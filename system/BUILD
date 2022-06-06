# A generic network operational interface gRPC service for interacting with
# network devices based on OpenConfig models. This is for the gNOI APIs used
# to perform diagnostic operations on a network device.

load("@com_github_grpc_grpc//bazel:cc_grpc_library.bzl", "cc_grpc_library")

package(default_visibility = ["//visibility:public"],
    licenses = ["notice"],)

filegroup(
    name = "system_proto_srcs",
    srcs = glob(["gnoi/common/*.proto",
                 "gnoi/types/*.proto"]),
)

proto_library(
    name = "system_proto",
    srcs = ["system.proto"],
    import_prefix = "github.com/openconfig/gnoi",
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:descriptor_proto",
        "//common:common_proto",
        "//types:types_proto",
    ],
)

cc_proto_library(
    name = "system_cc_proto",
    deps = [":system_proto"],
)

cc_grpc_library(
    name = "system_cc_grpc_proto",
    srcs = [":system_proto"],
    deps = [":system_cc_proto"],
    grpc_only = True,
)
