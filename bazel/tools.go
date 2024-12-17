//go:build tools

package bazel

// These imports only exist to keep go.mod entries for packages that are
// referenced in BUILD files, but not in Go code.

import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
