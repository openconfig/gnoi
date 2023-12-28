#!/bin/bash
set -euo pipefail

proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src/github.com/googleapis/googleapis:${GOPATH}/src:."

# Go
for p in types common containerz debug diag bgp cert factory_reset file healthz layer2 mpls system os otdr wavelength_router packet_link_qualification bootz; do
  protoc -I="${proto_imports}" --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative $p/$p.proto
done

