package main

// Imports and register the syslog-based query logger

import (
	_ "gopkg.in/sqle/vitess-go.v2/vt/vttablet/sysloglogger"
)
