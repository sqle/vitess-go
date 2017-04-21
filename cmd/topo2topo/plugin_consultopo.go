package main

// This plugin imports consultopo to register the consul implementation of TopoServer.

import (
	_ "gopkg.in/sqle/vitess-go.v2/vt/topo/consultopo"
)
