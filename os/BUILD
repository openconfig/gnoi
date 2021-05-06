#Copyright 2020 Google LLC
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Supporting infrastructure for implementing and testing PINS.
#

load("@com_github_grpc_grpc//bazel:cc_grpc_library.bzl", "cc_grpc_library")

package(
    default_visibility = ["//visibility:public"],
    licenses = ["notice"],
)

filegroup(
    name = "os_proto_srcs",
    srcs = glob(["**/*.proto"]),
    visibility = ["//third_party/pins_infra/thinkit:__pkg__"],
)

proto_library(
    name = "os_proto",
    srcs = ["os.proto"],
    import_prefix = "github.com/openconfig/gnoi",
    deps = [
        "//types:types_proto",
        "//common:common_proto",
    ],
)

cc_proto_library(
    name = "os_cc_proto",
    deps = [":os_proto"],
)

cc_grpc_library(
    name = "os_cc_grpc_proto",
    srcs = [":os_proto"],
    deps = [":os_cc_proto"],
    grpc_only = True,
)
