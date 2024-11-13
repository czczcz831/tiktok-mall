#!/bin/bash

svcName=${1}

echo ${svcName}.thrift
cd app/${svcName}
cwgo server  --type RPC --service ${svcName} --module github.com/czczcz831/tiktok-mall/app/${svcName} --idl ../../idl/${svcName}.thrift
cd ../../client/${svcName}
cwgo client  --type RPC --service ${svcName} --module github.com/czczcz831/tiktok-mall/client/${svcName} --idl ../../idl/${svcName}.thrift