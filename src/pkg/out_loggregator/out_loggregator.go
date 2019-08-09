package main

import (
	"code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"github.com/fluent/fluent-bit-go/output"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"fmt"
	"log"
	"net/http"
	"time"
)
import (
	"C"
	"unsafe"
)

var (
	client   *loggregator.IngressClient
	sourceId string

	ingress = promauto.NewCounter(prometheus.CounterOpts{
		Name: "service_logs_ingress",
		Help: "The total number of ingressed logs",
	})

	egress = promauto.NewCounter(prometheus.CounterOpts{
		Name: "service_logs_egress",
		Help: "The total number of egressed logs",
	})
)

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	return output.FLBPluginRegister(ctx, "loggregator", "Output to CF Loggregator")
}

//export FLBPluginInit
// (fluentbit will call this)
// ctx (context) pointer to fluentbit context (state/ c code)
func FLBPluginInit(ctx unsafe.Pointer) int {
	// Example to retrieve an optional configuration parameter
	cert := output.FLBPluginConfigKey(ctx, "cert")
	key := output.FLBPluginConfigKey(ctx, "key")
	ca := output.FLBPluginConfigKey(ctx, "ca")
	addr := output.FLBPluginConfigKey(ctx, "addr")
	sourceId = output.FLBPluginConfigKey(ctx, "source_id")
	port := output.FLBPluginConfigKey(ctx, "metrics_port")

	tlsConfig, err := loggregator.NewIngressTLSConfig(
		ca,
		cert,
		key,
	)
	if err != nil {
		log.Fatalf("Failed to create loggregator agent credentials: %s", err)
		return output.FLB_ERROR
	}

	client, err = loggregator.NewIngressClient(
		tlsConfig,
		loggregator.WithAddr(addr),
	)
	if err != nil {
		log.Fatalf("Failed to create loggregator agent client: %s", err)
		return output.FLB_ERROR
	}

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		if err != nil {
			log.Fatalf("Failed to start metrics server: %s", err)
		}
	}()

	return output.FLB_OK
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	var ret int
	var ts interface{}
	var record map[interface{}]interface{}

	// Create Fluent Bit decoder
	dec := output.NewDecoder(data, int(length))

	for {
		// Extract Record
		ret, ts, record = output.GetRecord(dec)
		if ret != 0 {
			break
		}
		ingress.Add(float64(len(record)))

		// Print record keys and values
		var timestamp time.Time
		switch tts := ts.(type) {
		case output.FLBTime:
			timestamp = tts.Time
		case uint64:
			// From our observation, when ts is of type uint64 it appears to
			// be the amount of seconds since unix epoch.
			timestamp = time.Unix(int64(tts), 0)
		default:
			timestamp = time.Now()
		}

		for _, v := range record {
			message := v.([]byte)
			e := &loggregator_v2.Envelope{
				SourceId:  sourceId,
				Timestamp: timestamp.UnixNano(),
				Message: &loggregator_v2.Envelope_Log{
					Log: &loggregator_v2.Log{
						Payload: message,
					},
				},
			}
			client.Emit(e)

			egress.Inc()
		}
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}
