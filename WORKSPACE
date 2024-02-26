load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

### Bazel rules for many languages to compile PROTO into gRPC libraries
http_archive(
    name = "rules_proto_grpc",
    sha256 = "c0d718f4d892c524025504e67a5bfe83360b3a982e654bc71fed7514eb8ac8ad",
    strip_prefix = "rules_proto_grpc-4.6.0",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/4.6.0.tar.gz"],
)

load(
    "@rules_proto_grpc//:repositories.bzl",
    "bazel_gazelle",
    "io_bazel_rules_go",
    "rules_proto_grpc_repos",
    "rules_proto_grpc_toolchains",
)

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

### Golang
io_bazel_rules_go()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(go_version = "1.20")

# gazelle:repo bazel_gazelle
bazel_gazelle()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_openconfig_bootz",
    importpath = "github.com/openconfig/bootz",
    sum = "h1:pCBi4GXcT+XM2Vud8kh0Q6uyw6oUxN1CFvrxnW+I/as=",
    version = "v0.1.1",
)

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

go_repository(
    name = "com_github_openconfig_gnsi",
    importpath = "github.com/openconfig/gnsi",
    sum = "h1:P6MjCnLZuINIivGLbp4No1HarZt7456wuJRNbSQyGu0=",
    version = "v1.2.5",
)

go_repository(
    name = "com_github_kylelemons_godebug",
    importpath = "github.com/kylelemons/godebug",
    sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_openconfig_goyang",
    importpath = "github.com/openconfig/goyang",
    sum = "h1:5MyIz4bN4vpH6aHDN339bkWXAjTkhg1ZKMhR4aIi5Rk=",
    version = "v0.0.0-20200115183954-d0a48929f0ea",
)

go_repository(
    name = "com_github_openconfig_ygot",
    build_directives = [
        "gazelle:proto_import_prefix github.com/openconfig/ygot",
    ],
    importpath = "github.com/openconfig/ygot",
    sum = "h1:kJJFPBrczC6TDnz/HMlFTJEdW2CuyUftV13XveIukg0=",
    version = "v0.6.0",
)

load("//:gnoi_deps.bzl", "go_dependnecies")

# gazelle:repository_macro gnoi_deps.bzl%go_dependnecies
go_dependnecies()

gazelle_dependencies()

load("@rules_proto_grpc//go:repositories.bzl", rules_proto_grpc_go_repos = "go_repos")

rules_proto_grpc_go_repos()

### C++
load("@rules_proto_grpc//cpp:repositories.bzl", rules_proto_grpc_cpp_repos = "cpp_repos")

rules_proto_grpc_cpp_repos()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

# open-config YANG files
http_archive(
    name = "github_openconfig_yang",
    build_file_content = """exports_files(glob(["release/models/**/*.yang"]), visibility = ["//visibility:public"])""",
    sha256 = "f6b2b6c0ffe0b66881287bcd43241a57583f353cc5cc41cba973601c32232f45",
    strip_prefix = "public-bf737a5567ec248456cb528efcd63cab15e8fc69",
    urls = [
        "https://github.com/openconfig/public/archive/bf737a5567ec248456cb528efcd63cab15e8fc69.zip",
    ],
)

# YANG files from other standard bodies.
http_archive(
    name = "github_yang",
    build_file_content = """exports_files(glob(["standard/**/*.yang"]), visibility = ["//visibility:public"])""",
    sha256 = "55913058f64a1ec7fe9e6e70d7128f08e66b20c859803b1fb02dbaf7eef2c64d",
    strip_prefix = "yang-2fa291d6bdb4b281d4e1b3dfa3254ffa7257d800",
    urls = [
        "https://github.com/YangModels/yang/archive/2fa291d6bdb4b281d4e1b3dfa3254ffa7257d800.zip",
    ],
)
