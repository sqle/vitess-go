package memorytopo

import (
	"testing"

	"gopkg.in/sqle/vitess-go.v2/vt/topo"
	"gopkg.in/sqle/vitess-go.v2/vt/topo/test"
)

func TestMemoryTopo(t *testing.T) {
	// Run the TopoServerTestSuite tests.
	test.TopoServerTestSuite(t, func() topo.Impl {
		return New("test")
	})
}
