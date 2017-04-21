package main

// This plugin imports ldapauthserver to register the LDAP implementation of AuthServer.

import (
	"gopkg.in/sqle/vitess-go.v2/mysqlconn/ldapauthserver"
	"gopkg.in/sqle/vitess-go.v2/vt/vtgate"
)

func init() {
	vtgate.RegisterPluginInitializer(func() { ldapauthserver.Init() })
}
