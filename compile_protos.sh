#!/bin/sh
set -euo pipefail

proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src"

# Go
protoc -I=$proto_imports --go_out=${GOPATH}/src types.proto
for p in bgp cert diag file interface layer2 mpls system; do
  protoc -I=$proto_imports --go_out=plugins=grpc:${GOPATH}/src $p/$p.proto
done
 
