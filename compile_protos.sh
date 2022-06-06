#!/bin/bash
set -euo pipefail

proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src"

# Go
for p in types common diag bgp cert file healthz interface layer2 mpls system otdr wavelength_router; do
  protoc -I="${proto_imports}" --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative $p/$p.proto
done

