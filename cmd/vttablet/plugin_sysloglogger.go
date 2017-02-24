package main

// Imports and register the syslog-based query logger

import (
	_ "gopkg.in/sqle/vitess-go.v1/vt/tabletserver/sysloglogger"
)
