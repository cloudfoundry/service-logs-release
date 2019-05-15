## Service Logs Release

This release contains a job that allows you to send logs from syslog files to Loggregator. The
job can be added to a service VM and configured to forward that service's logs to Loggregator.

## How it works
Service Logs Release uses the [Fluent Bit Loggregator Plugin][fluent-bit-loggr-plugin] to read from syslog 
files and send those logs to Loggregator. The Fluent Bit Loggregator Plugin is a plugin for [Fluent Bit][fluent-bit].

## Configuration
The sample ops file (`manifests/operations/add-service-logs.yml`) demonstrates the service logs job on the api VM.

[fluent-bit-loggr-plugin]: https://github.com/cloudfoundry/service-logs-release/blob/develop/src/pkg/out_loggregator/out_loggregator.go
[fluent-bit]: https://fluentbit.io/
