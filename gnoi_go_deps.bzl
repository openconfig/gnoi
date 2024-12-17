load("@bazel_gazelle//:deps.bzl", "go_repository")

def gnoi_go_deps():
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:G2HAfAmvm/GcKan2oOQpBXOd2tT2G57ZnZGWa1PxPBQ=",
        version = "v4.1.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:UL815xU9SqsFlibzuggzjXhog7bL6oX9BbNZnL2UFvs=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_chappjc_logrus_prefix",
        importpath = "github.com/chappjc/logrus-prefix",
        sum = "h1:aZTKxMminKeQWHtzJBbV8TttfTxzdJ+7iEJFE6FmUzg=",
        version = "v0.0.0-20180227015900-3a1d64819adb",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:N+3sFI5GUjRKBi+i0TxYVST9h4Ie192jJWpHvthBBgg=",
        version = "v0.0.0-20240723142845-024c85f92f20",
    )
    go_repository(
        name = "com_github_coredhcp_coredhcp",
        importpath = "github.com/coredhcp/coredhcp",
        sum = "h1:8bGxpjqPoic463Lr0Uqnnz7P/xSZCPKDkxkjeywS79k=",
        version = "v0.0.0-20230808195049-3e32ddb5ac86",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:HzkeUz1Knt+3bK+8LG1bxOO/jzWZmdxpwC51i202les=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:tntQDh69XqOCOZsDz0lVJQez/2L6Uu2PdjCQwWCJ3bM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:n+5WquG0fcWoWp6xPWfHdbskMCQaFnG6PfBrh1Ky4HY=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:1+mZ9upx1Dh6FmUTFR1naJ77miKiXgALjWOZ3NVFPmY=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_gopacket",
        importpath = "github.com/google/gopacket",
        sum = "h1:ves8RnFZPGiFnTS0uPQStjwru6uO6h+nlr9j6fL7kF8=",
        version = "v1.1.19",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_h_fam_errdiff",
        importpath = "github.com/h-fam/errdiff",
        sum = "h1:rPsW4ob2fMOIulwTEoZXaaUIuud7XUudw5SLKTZj3Ss=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_insomniacslk_dhcp",
        importpath = "github.com/insomniacslk/dhcp",
        sum = "h1:S33o3djA1nPRd+d/bf7jbbXytXuK/EoXow7+aa76grQ=",
        version = "v0.0.0-20230908212754-65c27093e38a",
    )
    go_repository(
        name = "com_github_josharian_native",
        importpath = "github.com/josharian/native",
        sum = "h1:uuaP0hAbW7Y4l0ZRQ6C9zfb7Mg1mbFKry/xzDAfmtLA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:JITubQf0MOLdlGRuRq+jtsDlekdYPia9ZFsB8h/APPA=",
        version = "v0.0.19",
    )
    go_repository(
        name = "com_github_mgutz_ansi",
        importpath = "github.com/mgutz/ansi",
        sum = "h1:5PJl274Y63IEHC+7izoQE9x6ikvDFZS2mDVS3drnohI=",
        version = "v0.0.0-20200706080929-d51e80ef957d",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_nxadm_tail",
        importpath = "github.com/nxadm/tail",
        sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_openconfig_bootz",
        importpath = "github.com/openconfig/bootz",
        sum = "h1:Q0mThGmZiX/kht+crar6FtLVxqdjUS/deMnpcNX+F7c=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_openconfig_gnmi",
        build_directives = [
            "gazelle:proto_import_prefix github.com/openconfig/gnmi",
        ],
        build_file_generation = "on",
        importpath = "github.com/openconfig/gnmi",
        sum = "h1:H7pLIb/o3xObu3+x0Fv9DCK7TH3FUh7mNwbYe+34hFw=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_github_openconfig_gnsi",
        importpath = "github.com/openconfig/gnsi",
        sum = "h1:PfQa9Gy0lH1sHqA2L3Gj2fEh2zPMbWxMmIRQv2Nk1T8=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_openconfig_goyang",
        importpath = "github.com/openconfig/goyang",
        sum = "h1:5MyIz4bN4vpH6aHDN339bkWXAjTkhg1ZKMhR4aIi5Rk=",
        version = "v0.0.0-20200115183954-d0a48929f0ea",
    )
    go_repository(
        name = "com_github_openconfig_grpctunnel",
        importpath = "github.com/openconfig/grpctunnel",
        sum = "h1:t6SvvdfWCMlw0XPlsdxO8EgO+q/fXnTevDjdYREKFwU=",
        version = "v0.0.0-20220819142823-6f5422b8ca70",
    )
    go_repository(
        name = "com_github_openconfig_ygot",
        importpath = "github.com/openconfig/ygot",
        sum = "h1:kJJFPBrczC6TDnz/HMlFTJEdW2CuyUftV13XveIukg0=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:uH2qQXheeefCCkuBBSLi7jCiSmj3VRh2+Goq2N7Xxu0=",
        version = "v2.0.9",
    )
    go_repository(
        name = "com_github_pierrec_lz4_v4",
        importpath = "github.com/pierrec/lz4/v4",
        sum = "h1:xaKrnTkyoqfh1YItXl56+6KJNVYWlEEPuAQW9xsplYQ=",
        version = "v4.1.18",
    )
    go_repository(
        name = "com_github_planetscale_vtprotobuf",
        importpath = "github.com/planetscale/vtprotobuf",
        sum = "h1:GFCKgmp0tecUJ0sJuv4pzYCqS9+RGSn52M3FUwPs+uo=",
        version = "v0.6.1-0.20240319094008-0393e58bdf10",
    )
    go_repository(
        name = "com_github_protocolbuffers_txtpbfmt",
        importpath = "github.com/protocolbuffers/txtpbfmt",
        sum = "h1:AKJY61V2SQtJ2a2PdeswKk0NM1qF77X+julRNYRxPOk=",
        version = "v0.0.0-20220608084003-fc78c767cd6a",
    )
    go_repository(
        name = "com_github_rifflock_lfshook",
        importpath = "github.com/rifflock/lfshook",
        sum = "h1:mZHayPoR0lNmnHyvtYjDeq0zlVHn9K/ZXoy17ylucdo=",
        version = "v0.0.0-20180920164130-b9218ef580f5",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:stMpOSZFs//0Lv29HduCmli3GUfpFoF3Y1Q/aXj/wVM=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:R+kOtfhWQE6TVQzY+4D7wJLBgkdVasCEFxSUBYBYIlA=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:zqmyTlQyufRC65JnImJ6H1Sf7BDj8bG31EV919NVEQc=",
        version = "v1.0.6-0.20201009195203-85dd5c8bc61c",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        sum = "h1:rGGH0XDZhdUOryiDWjmIvUSWpbNqisK8Wk0Vyefw8hc=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:X1TuBLAMDFbaTAChgCBLu3DU3UPyELpnF2jjJ2cz/S8=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_u_root_uio",
        importpath = "github.com/u-root/uio",
        sum = "h1:YcojQL98T/OO+rybuzn2+5KrD5dBwXIvYBvQ2cD3Avg=",
        version = "v0.0.0-20230305220412-3e8cd9d6bf63",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:Zr0eK8JbFv6+Wi4ilXAR8FJ3wyNdpxHKJNPos6LTZOY=",
        version = "v0.5.0",
    )
    go_repository(
        name = "dev_cel_expr",
        importpath = "cel.dev/expr",
        sum = "h1:yloc84fytn4zmJX2GU3TkXGsaieaV7dQ057Qs4sIG2Y=",
        version = "v0.16.0",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
        version = "v1.67.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:KpwkzHKEF7B9Zxg18WzOa7djJ+Ha5DzthMyZYQfEn2A=",
        version = "v0.0.0-20230410155749-daa745c078e1",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:wKguEg1hsxI2/L3hUYrpo1RVi48K+uTyzKqprwLXsb8=",
        version = "v0.0.0-20240814211410-ddb44dafa142",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:e7S5W7MGGLaSu8j3YjdezkZ+m1/Nm0uRVRMEMGk26Xs=",
        version = "v0.0.0-20240814211410-ddb44dafa142",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:zWnc1Vrcno+lHZCOofnIMvycFcc0QRGIzm9dhnDX68E=",
        version = "v1.67.1",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        build_file_proto_mode = "disable_global",  # Manually added to fix build. See https://github.com/golang/protobuf/issues/1611
        importpath = "google.golang.org/protobuf",
        sum = "h1:m3LfL6/Ca+fqnjnlqQXNpFPABW1UD7mjh8KO2mKFytA=",
        version = "v1.35.1",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:RrRspgV4mU+YwB4FYnuBoKsUapNIL5cohGAmSH3azsw=",
        version = "v0.26.0",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:zY54UmvipHiNd+pm+m0x9KhZ9hl1/7QNMyxXbc6ICqA=",
        version = "v0.17.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:a9JDOJc5GMUJ0+UDqmLT86WiEy7iWyIhz8gz8E4e5hE=",
        version = "v0.28.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:BzDx2FehcG7jJwgWLELCdmLuxk2i+x9UDpSiss2u0ZA=",
        version = "v0.22.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:3NFvSEYkUoMifnESzZl15y791HH1qU2xm6eCJU5ZPXQ=",
        version = "v0.8.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:Twjiwq9dn6R1fQcyiK+wQyHWfaz/BJB+YIpzU/Cv3Xg=",
        version = "v0.24.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:F6D4vR+EHoL9/sWAWgAR1H2DcHr4PareCbAaCo1RpuU=",
        version = "v0.23.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:XtiM5bkSOt+ewxlOE/aE/AKEHibwj/6gvWMl9Rsh0Qc=",
        version = "v0.17.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:vU5i/LfpvrRCpgM/VPfJLg5KjxD3E+hfT1SH+d9zLwg=",
        version = "v0.21.1-0.20240508182429-e35e4ccd0d2d",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
        version = "v0.0.0-20200804184101-5ec99f83aff1",
    )
    go_repository(
        name = "org_mozilla_go_pkcs7",
        importpath = "go.mozilla.org/pkcs7",
        sum = "h1:CCriYyAfq1Br1aIYettdHZTy8mBTIPo7We18TuO/bak=",
        version = "v0.0.0-20210826202110-33d05740a352",
    )
