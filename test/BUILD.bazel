load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_test(
    name = "simple_test",
    size = "small",
    srcs = ["simple_test.go"],
    deps = [
        "//github.com/golang/protobuf/proto:proto",
        "//github.com/openconfig/reference/rpc/gnoi:go_default_library",
        "//github.com/openconfig/reference/rpc/gnoi/bgp:go_default_library",
        "//github.com/openconfig/reference/rpc/gnoi/cert:go_default_library",
        "//github.com/openconfig/reference/rpc/gnoi/file:go_default_library",
        "//github.com/openconfig/reference/rpc/gnoi/interface:go_default_library",
        "//github.com/openconfig/reference/rpc/gnoi/layer2:go_default_library",
        "//github.com/openconfig/reference/rpc/gnoi/mpls:go_default_library",
        "//github.com/openconfig/reference/rpc/gnoi/system:go_default_library",
    ],
)
