#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/auth-svc"
exec "$CURDIR/bin/auth-svc"
