# gNOI - gRPC Network Operations Interface
gNOI defines a set of gRPC-based microservices for executing operational commands on network devices.

# Rebuild *.pb.go files
protoc -I=${GOPATH}/src:${GOPATH}/src/github.com/openconfig/gnoi --go_out=plugins=grpc:.  ${GOPATH}/src/github.com/openconfig/gnoi/types.proto
protoc -I=${GOPATH}/src:${GOPATH}/src/github.com/openconfig/gnoi --go_out=plugins=grpc:.  ${GOPATH}/src/github.com/openconfig/gnoi/bgp/bgp.proto
protoc -I=${GOPATH}/src:${GOPATH}/src/github.com/openconfig/gnoi --go_out=plugins=grpc:.  ${GOPATH}/src/github.com/openconfig/gnoi/cert/cert.proto
protoc -I=${GOPATH}/src:${GOPATH}/src/github.com/openconfig/gnoi --go_out=plugins=grpc:.  ${GOPATH}/src/github.com/openconfig/gnoi/interface/interface.proto
protoc -I=${GOPATH}/src:${GOPATH}/src/github.com/openconfig/gnoi --go_out=plugins=grpc:.  ${GOPATH}/src/github.com/openconfig/gnoi/layer2/layer2.proto
protoc -I=${GOPATH}/src:${GOPATH}/src/github.com/openconfig/gnoi --go_out=plugins=grpc:.  ${GOPATH}/src/github.com/openconfig/gnoi/mpls/mpls.proto
protoc -I=${GOPATH}/src:${GOPATH}/src/github.com/openconfig/gnoi --go_out=plugins=grpc:.  ${GOPATH}/src/github.com/openconfig/gnoi/system/system.proto
