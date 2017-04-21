package testsuite

import (
	"testing"

	"gopkg.in/sqle/vitess-go.v2/vt/callerid"
	querypb "gopkg.in/sqle/vitess-go.v2/vt/proto/query"
	vtrpcpb "gopkg.in/sqle/vitess-go.v2/vt/proto/vtrpc"
)

func TestFakeCallerID(t *testing.T) {
	im := querypb.VTGateCallerID{
		Username: FakeUsername,
	}
	ef := vtrpcpb.CallerID{
		Principal:    FakePrincipal,
		Component:    FakeComponent,
		Subcomponent: FakeSubcomponent,
	}
	RunTests(t, &im, &ef, callerid.NewContext)
}
