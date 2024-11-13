#!/bin/bash

cd app/api
echo api.thrift
cwgo server --type http --service api --module github.com/czczcz831/tiktok-mall/app/api --idl ../../idl/api.thrift
cwgo client --type http --server_name localhost --module github.com/czczcz831/tiktok-mall/app/api  --idl ../../idl/api.thrift