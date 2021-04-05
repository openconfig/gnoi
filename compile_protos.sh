#!/bin/sh
set -euo pipefail

proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src"

# Go
for p in types common diag bgp cert file interface layer2 mpls system os otdr wavelength_router; do
  protoc -I=$proto_imports --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $p/$p.proto
done

