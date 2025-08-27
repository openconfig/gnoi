# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
"""Dependencies to build gnoi."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def gnoi_deps():
    """Declare the third-party dependencies necessary to build gnoi"""
    if not native.existing_rule("bazel_features"):
        http_archive(
            name = "bazel_features",
            sha256 = "06a9691becc357c1e4af011b9fb3239097503ce00607d9eb6b11ea0fac304039",
            strip_prefix = "bazel_features-1.35.0",
            url = "https://github.com/bazel-contrib/bazel_features/releases/download/v1.35.0/bazel_features-v1.35.0.tar.gz",
        )
    if not native.existing_rule("bazel_gazelle"):
        http_archive(
            name = "bazel_gazelle",
            sha256 = "e467b801046b6598c657309b45d2426dc03513777bd1092af2c62eebf990aca5",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.45.0/bazel-gazelle-v0.45.0.tar.gz",
                "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.45.0/bazel-gazelle-v0.45.0.tar.gz",
            ],
        )
    if not native.existing_rule("com_github_grpc_grpc"):
        http_archive(
            name = "com_github_grpc_grpc",
            url = "https://github.com/grpc/grpc/archive/refs/tags/v1.74.1.tar.gz",
            strip_prefix = "grpc-1.74.1",
            sha256 = "7bf97c11cf3808d650a3a025bbf9c5f922c844a590826285067765dfd055d228",
        )
    if not native.existing_rule("com_google_googleapis"):
        http_archive(
            name = "com_google_googleapis",
            sha256 = "3e8a622f6e72e1660c16ae846af6dab0b3eb696fbdf0172e57b9cdd6752378e7",
            strip_prefix = "googleapis-fe8ba054ad4f7eca946c2d14a63c3f07c0b586a0",
            urls = ["https://github.com/googleapis/googleapis/archive/fe8ba054ad4f7eca946c2d14a63c3f07c0b586a0.zip"],
        )
    if not native.existing_rule("com_google_protobuf"):
        http_archive(
            name = "com_google_protobuf",
            url = "https://github.com/protocolbuffers/protobuf/archive/refs/tags/v32.0.zip",
            strip_prefix = "protobuf-32.0",
            sha256 = "3e8bdbefc2a9ef70c530d7ab20bfce6f4a10e5f9e101bacc756e479ab6c7de4f",
            repo_mapping = {
                "@proto_bazel_features": "@bazel_features",
            },
        )
    if not native.existing_rule("bazel_skylib"):
        http_archive(
            name = "bazel_skylib",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.8.1/bazel-skylib-1.8.1.tar.gz",
                "https://github.com/bazelbuild/bazel-skylib/releases/download/1.8.1/bazel-skylib-1.8.1.tar.gz",
            ],
            sha256 = "51b5105a760b353773f904d2bbc5e664d0987fbaf22265164de65d43e910d8ac",
        )
    if not native.existing_rule("io_bazel_rules_go"):
        http_archive(
            name = "io_bazel_rules_go",
            sha256 = "a729c8ed2447c90fe140077689079ca0acfb7580ec41637f312d650ce9d93d96",
            urls = [
                "https://github.com/bazelbuild/rules_go/releases/download/v0.57.0/rules_go-v0.57.0.zip",
            ],
        )
    if not native.existing_rule("rules_proto"):
        http_archive(
            name = "rules_proto",
            sha256 = "14a225870ab4e91869652cfd69ef2028277fc1dc4910d65d353b62d6e0ae21f4",
            strip_prefix = "rules_proto-7.1.0",
            url = "https://github.com/bazelbuild/rules_proto/releases/download/7.1.0/rules_proto-7.1.0.tar.gz",
        )
    if not native.existing_rule("openconfig_bootz"):
        http_archive(
            name = "openconfig_bootz",
            sha256 = "c8af009b7a38c83f44103889b487c3c36b5d63803c3b2fcda56dce977f115e80",
            strip_prefix = "bootz-028126f868486197a9c0bbc2a94810b1bbb3cf27",
            url = "https://github.com/bstoll/bootz/archive/028126f868486197a9c0bbc2a94810b1bbb3cf27.zip",
        )
    if not native.existing_rule("openconfig_gnmi"):
        http_archive(
            name = "openconfig_gnmi",
            sha256 = "813f8a52dfa06dd1b9a2c775b26c42d36a05595dfa6fb0a85dbaead46b5c43a3",
            strip_prefix = "gnmi-0.14.1",
            url = "https://github.com/openconfig/gnmi/archive/refs/tags/v0.14.1.tar.gz",
        )
    if not native.existing_rule("openconfig_gnsi"):
        http_archive(
            name = "openconfig_gnsi",
            sha256 = "df4c69885b14bb5c69a90dc4f9c0cfb78a6638a6404a79d70553d14fe350404a",
            strip_prefix = "gnsi-1.9.0",
            url = "https://github.com/openconfig/gnsi/archive/refs/tags/v1.9.0.tar.gz",
        )
