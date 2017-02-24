package vtctl

import (
	// This plugin imports zk2topo to register the zk2 implementation of TopoServer.
	_ "gopkg.in/sqle/vitess-go.v1/vt/topo/zk2topo"
)
