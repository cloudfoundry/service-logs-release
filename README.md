## Service Logs Release

This release adds a job that allows you to send logs from a syslog file to Loggregator

### Building the release
1. `git clone https://github.com/cloudfoundry/service-logs-release`
1. `cd service-logs-release`
1. `git submodule update --init`
1. `./src/code.cloudfoundry.org/fluentbit-loggr-plugin/generate_binaries.sh`
1. `bosh create-release`
