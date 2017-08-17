package simple

import (
	"testing"

	"github.com/golang/protobuf/proto/proto"
	"github.com/openconfig/gnoi"
	gbgp "github.com/openconfig/gnoi/bgp"
)

func TestGNOI(t *testing.T) {
	tests := []struct {
		desc string
		in   proto.Message
		want string
	}{{
		desc: "gnoi.Path",
		in: &gnoi.Path{
			Origin: "oc",
			Elem:   []*gnoi.PathElem{{Name: "interfaces", Key: map[string]string{"name": "Ethernet1/1/0"}}},
		},
		want: "origin: \"oc\"\nelem: <\n  name: \"interfaces\"\n  key: <\n    key: \"name\"\n    value: \"Ethernet1/1/0\"\n  >\n>\n",
	}, {
		desc: "gnoi.HashType",
		in: &gnoi.HashType{
			Method: gnoi.HashType_MD5,
			Hash:   []byte("foo"),
		},
		want: "method: MD5\nhash: \"foo\"\n",
	}, {
		desc: "bgp.ClearBGPNeighborRequest",
		in: &gbgp.ClearBGPNeighborRequest{
			Address:         "foo",
			RoutingInstance: "bar",
			Mode:            gbgp.ClearBGPNeighborRequest_HARD,
		},
		want: "address: \"foo\"\nrouting_instance: \"bar\"\nmode: HARD\n",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := proto.MarshalTextString(tt.in); got != tt.want {
				t.Fatalf("Text Proto output failed: got %q, want %q", got, tt.want)
			}
		})
	}
}
