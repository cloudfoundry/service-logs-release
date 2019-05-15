## Service Logs Release

This release contains a job that allows you to send logs from syslog files to Loggregator. This
job can be added to a service VM and configured to forward that service's logs to Loggregator.

## How it works
Service Logs Release uses the [Fluent Bit Loggregator Plugin][fluent-bit-loggr-plugin] to read from syslog 
files and send those logs to Loggregator. The Fluent Bit Loggregator Plugin is a plugin for [Fluent Bit][fluent-bit]

## Configuration
The sample ops file (`manifests/operations/add-service-logs.yml`) adds the service logs job to the `api` VM
and can be changed to add the service logs job to any VM.

### Building the release
1. `git clone https://github.com/cloudfoundry/service-logs-release`
1. `cd service-logs-release`
1. `git submodule update --init`
1. `./src/code.cloudfoundry.org/fluentbit-loggr-plugin/generate_binaries.sh`
1. `bosh create-release`

[fluent-bit-loggr-plugin]: https://github.com/cloudfoundry/fluentbit-loggr-plugin/tree/a6dbe6032bd4d24c026ac244a9993c4328869575
[fluent-bit]: https://fluentbit.io/