#!/bin/sh
set -euo pipefail

proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src"

# Go
for p in types common diag bgp cert file interface layer2 mpls system; do
  protoc -I=$proto_imports --go_out=plugins=grpc:. $p/$p.proto
done
 
