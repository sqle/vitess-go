package main

// This plugin imports consultopo to register the consul implementation of TopoServer.

import (
	_ "gopkg.in/sqle/vitess-go.v1/vt/topo/consultopo"
)
