package vtctl

import (
	// Imports etcd2topo to register the etcd2 implementation of
	// TopoServer.
	_ "gopkg.in/sqle/vitess-go.v2/vt/topo/etcd2topo"
)
