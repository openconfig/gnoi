#!/bin/bash
set -euo pipefail

if [ ! -z "${GOPATH:-}"]; then
proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src/github.com/googleapis/googleapis:${GOPATH}/src:."
else 
go get github.com/googleapis/googleapis
genproto=$(go list -m -f '{{.Dir}}' github.com/googleapis/googleapis)
proto_imports=".:${genproto}"
fi

# Go
for p in types common diag bgp cert factory_reset file healthz layer2 mpls system os otdr wavelength_router packet_link_qualification; do
  protoc -I="${proto_imports}" --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative $p/$p.proto
done

go mod tidy