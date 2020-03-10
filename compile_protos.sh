#!/bin/sh
set -euo pipefail

proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src"

declare -a pp=(types common diag bgp cert file interface layer2 mpls system otdr wavelength_router os factory_reset)

# Go
for p in "${pp[@]}"; do
  protoc -I=$proto_imports --go_out=plugins=grpc:. $p/$p.proto
done

