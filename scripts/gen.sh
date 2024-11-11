#!/bin/bash

svcName=${1}

cd app/${svcName}
echo ../idl/${svcName}.proto
cwgo client -I ../../idl --type RPC --service ${svcName} --module github.com/czczcz831/tiktok-mall/app/${svcName} --idl ../../idl/${svcName}.proto
cwgo server -I ../../idl --type RPC --service ${svcName} --module github.com/czczcz831/tiktok-mall/app/${svcName} --idl ../../idl/${svcName}.proto