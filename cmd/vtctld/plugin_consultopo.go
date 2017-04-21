package main

// Imports and register the 'consul' topo.Server and its Explorer.

import (
	"gopkg.in/sqle/vitess-go.v2/vt/servenv"
	"gopkg.in/sqle/vitess-go.v2/vt/topo/consultopo"
	"gopkg.in/sqle/vitess-go.v2/vt/vtctld"
)

func init() {
	// Wait until flags are parsed, so we can check which topo server is in use.
	servenv.OnRun(func() {
		if s, ok := ts.Impl.(*consultopo.Server); ok {
			vtctld.HandleExplorer("consul", vtctld.NewBackendExplorer(s))
		}
	})
}
