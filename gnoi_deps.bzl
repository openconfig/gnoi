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
            sha256 = "8b1c9b7558498000f5adebbc584b7bf15b6b2bf181448a66f6b2fc5b4c84231c",
            strip_prefix = "bazel_features-1.23.0",
            url = "https://github.com/bazel-contrib/bazel_features/releases/download/v1.23.0/bazel_features-v1.23.0.tar.gz",
        )
    if not native.existing_rule("bazel_gazelle"):
        http_archive(
            name = "bazel_gazelle",
            sha256 = "aefbf2fc7c7616c9ed73aa3d51c77100724d5b3ce66cfa16406e8c13e87c8b52",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.41.0/bazel-gazelle-v0.41.0.tar.gz",
                "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.41.0/bazel-gazelle-v0.41.0.tar.gz",
            ],
        )
    if not native.existing_rule("com_github_grpc_grpc"):
        http_archive(
            name = "com_github_grpc_grpc",
            url = "https://github.com/grpc/grpc/archive/refs/tags/v1.73.1.tar.gz",
            strip_prefix = "grpc-1.73.1",
            sha256 = "e11fd9b963c617de53d08a84f41359164b123f2e8e4180644706688fc9de43d9",
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
            url = "https://github.com/protocolbuffers/protobuf/archive/refs/tags/v31.1.zip",
            strip_prefix = "protobuf-31.1",
            sha256 = "fc6289aa4450bdb70626aceaaebebdd7d3d4725c288a9cbb138a26defe5d9987",
            repo_mapping = {
                "@proto_bazel_features": "@bazel_features",
            },
        )
    if not native.existing_rule("bazel_skylib"):
        http_archive(
            name = "bazel_skylib",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.7.1/bazel-skylib-1.7.1.tar.gz",
                "https://github.com/bazelbuild/bazel-skylib/releases/download/1.7.1/bazel-skylib-1.7.1.tar.gz",
            ],
            sha256 = "bc283cdfcd526a52c3201279cda4bc298652efa898b10b4db0837dc51652756f",
        )
    if not native.existing_rule("io_bazel_rules_go"):
        http_archive(
            name = "io_bazel_rules_go",
            sha256 = "0936c9bc3c4321ee372cb8f66dd972d368cb940ed01a9ba9fd7debcf0093f09b",
            urls = [
                "https://github.com/bazelbuild/rules_go/releases/download/v0.51.0/rules_go-v0.51.0.zip",
            ],
        )
    if not native.existing_rule("rules_proto"):
        http_archive(
            name = "rules_proto",
            sha256 = "0e5c64a2599a6e26c6a03d6162242d231ecc0de219534c38cb4402171def21e8",
            strip_prefix = "rules_proto-7.0.2",
            url = "https://github.com/bazelbuild/rules_proto/releases/download/7.0.2/rules_proto-7.0.2.tar.gz",
        )
    if not native.existing_rule("openconfig_bootz"):
        http_archive(
            name = "openconfig_bootz",
            sha256 = "e47b53b9eebee49b69df735d9b65ec9da00bca5c766910ec2b4abc62c0006915",
            strip_prefix = "bootz-dac82dd3345fe8c0d2759fbe63341448a58e7a7d",
            url = "https://github.com/bstoll/bootz/archive/dac82dd3345fe8c0d2759fbe63341448a58e7a7d.zip",
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
