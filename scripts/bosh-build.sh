#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source $DIR/shared

echo "-- BOSH create-release"
prep_src
create_release release.tgz

boshdir=$(mktemp -d)
echo "-- BOSH extract and compile in ${boshdir}"
pushd ${boshdir}
  tar -xzf ${DIR}/../release.tgz

  pushd packages
    mkdir -p loggregator-output-plugin
    tar -xzf loggregator-output-plugin.tgz -C loggregator-output-plugin

    pushd loggregator-output-plugin
      go build -buildmode=c-shared -mod=vendor -o ./out_loggregator.so ./src/pkg/out_loggregator
    popd
  popd
popd

cleanup
rm -f ${DIR}/../release.tgz
rm -rf ${boshdir}