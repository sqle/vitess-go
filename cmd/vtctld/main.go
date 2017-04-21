package main

import (
	"flag"

	"gopkg.in/sqle/vitess-go.v2/vt/servenv"
	"gopkg.in/sqle/vitess-go.v2/vt/topo"
	"gopkg.in/sqle/vitess-go.v2/vt/vtctld"
)

func init() {
	servenv.RegisterDefaultFlags()
}

// used at runtime by plug-ins
var (
	ts topo.Server
)

func main() {
	flag.Parse()
	servenv.Init()
	defer servenv.Close()

	ts = topo.Open()
	defer ts.Close()

	// Init the vtctld core
	vtctld.InitVtctld(ts)

	// Start schema manager service.
	initSchema()

	// And run the server.
	servenv.RunDefault()
}
