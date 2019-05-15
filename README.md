## Service Logs Release

This release contains a job that allows you to send logs from syslog files to Loggregator. This
job can be added to a service VM and configured to forward that service's logs to Loggregator.

## How it works
Service Logs Release uses the [Fluent Bit Loggregator Plugin][fluent-bit-loggr-plugin] to read from syslog 
files and send those logs to Loggregator. The Fluent Bit Loggregator Plugin is a plugin for [Fluent Bit][fluent-bit].

## Configuration

 Service Logs release requires the following configuration properties:
   
   | property name | description                                                          |
   | ------------- | -------------------------------------------------------------------- |
   | source_id     | Source id of the logs to be emitted                                  |
   | path          | File glob representing the logs to be emitted                        |
   | cert          | Cert for cert-key pair that can connect to a local loggregator agent |
   | key           | Key for cert-key pair that can connect to a local loggregator agent  |
   | ca            | CA for cert-key pair that can connect to a local loggregator agent   |
   | port          | Port where the loggregator agent is listening                        |


The sample ops file (`manifests/operations/add-service-logs.yml`) demonstrates a sample configuration that adds the service logs job to the `api`.

### Building the release
1. `git clone https://github.com/cloudfoundry/service-logs-release`
1. `cd service-logs-release`
1. `./scripts/bosh-build.sh`

[fluent-bit-loggr-plugin]: https://github.com/cloudfoundry/fluentbit-loggr-plugin
[fluent-bit]: https://fluentbit.io/
