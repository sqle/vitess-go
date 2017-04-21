package main

// This plugin imports staticauthserver to register the flat-file implementation of AuthServer.

import (
	"gopkg.in/sqle/vitess-go.v2/mysqlconn"
	"gopkg.in/sqle/vitess-go.v2/vt/vtgate"
)

func init() {
	vtgate.RegisterPluginInitializer(func() { mysqlconn.InitAuthServerStatic() })
}
