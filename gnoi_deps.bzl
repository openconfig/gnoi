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
load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def gnoi_deps():
    """Declare the third-party dependencies necessary to build gnoi"""
    if not native.existing_rule("rules_proto_grpc"):
      http_archive(
        name = "rules_proto_grpc",
        sha256 = "c0d718f4d892c524025504e67a5bfe83360b3a982e654bc71fed7514eb8ac8ad",
        strip_prefix = "rules_proto_grpc-4.6.0",
        urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/4.6.0.tar.gz"],
      )
    if not native.existing_rule("com_github_grpc_grpc"):
      http_archive(
        name = "com_github_grpc_grpc",
        url = "https://github.com/grpc/grpc/archive/refs/tags/v1.61.1.tar.gz",
        strip_prefix = "grpc-1.61.1",
        sha256 = "b74ce7d26fe187970d1d8e2c06a5d3391122f7bc1fdce569aff5e435fb8fe780",
      )
    if not native.existing_rule("com_google_protobuf"):
      http_archive(
        name = "com_google_protobuf",
        url = "https://github.com/protocolbuffers/protobuf/releases/download/v4.25.2/protobuf-all-25.2.tar.gz",
        strip_prefix = "protobuf-",
        sha256 = "ba0650be1b169d24908eeddbe6107f011d8df0da5b1a5a4449a913b10e578faf",
      )
    if not native.existing_rule("com_github_openconfig_bootz"):
        go_repository(
            name = "com_github_openconfig_bootz",
            importpath = "github.com/openconfig/bootz",
            sum = "h1:JTeVs+DCy3pjc4X1sifuBEqeNgRK7m7OL8gdDgHfplA=",
            version = "v0.2.1",
        )
    if not native.existing_rule("com_github_openconfig_gnmi"):
        go_repository(
            name = "com_github_openconfig_gnmi",
            build_directives = [
                "gazelle:proto_import_prefix github.com/openconfig/gnmi",
            ],    
            build_file_generation = "on",
            importpath = "github.com/openconfig/gnmi",
            sum = "h1:tv9HygDMXnoGyWuLmNCodMV2+PK6+uT/ndAxDVzsUUQ=",
            version = "v0.0.0-20220617175856-41246b1b3507",
        )
    if not native.existing_rule("com_github_openconfig_gnsi"):
        go_repository(
            name = "com_github_openconfig_gnsi",
            importpath = "github.com/openconfig/gnsi",
            sum = "h1:P6MjCnLZuINIivGLbp4No1HarZt7456wuJRNbSQyGu0=",
            version = "v1.2.5",
        )
    if not native.existing_rule("com_github_kylelemons_godebug"):
        go_repository(
            name = "com_github_kylelemons_godebug",
            importpath = "github.com/kylelemons/godebug",
            sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
            version = "v1.1.0",
        )
    if not native.existing_rule("com_github_openconfig_goyang"):
        go_repository(
            name = "com_github_openconfig_goyang",
            importpath = "github.com/openconfig/goyang",
            sum = "h1:Z95LskKYk6nBYOxHtmJCu3YEKlr3pJLWG1tYAaNh3yU=",
            version = "v0.2.9",
        )
    if not native.existing_rule("com_github_openconfig_ygot"):
        go_repository(
            name = "com_github_openconfig_ygot",
            build_directives = [
                "gazelle:proto_import_prefix github.com/openconfig/ygot",
            ],
            importpath = "github.com/openconfig/ygot",
            sum = "h1:EKaeFhx1WwTZGsYeqipyh1mfF8y+z2StaXZtwVnXklk=",
            version = "v0.13.1",
        )
