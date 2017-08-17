package simple

import (
	"testing"

	"github.com/golang/protobuf/proto/proto"

	gnoi "github.com/openconfig/reference/rpc/gnoi"
	gbgp "github.com/openconfig/reference/rpc/gnoi/bgp"
)

func TestGNOI(t *testing.T) {
	tests := []struct {
		desc string
		in   proto.Message
		want string
	}{{
		desc: "gnoi",
		in: &gnoi.Path{
			Elements: []string{"foo", "path"},
		},
		want: "elements: \"foo\"\nelements: \"path\"\n",
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
