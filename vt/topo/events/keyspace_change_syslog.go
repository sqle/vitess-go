package events

import (
	"fmt"
	"log/syslog"

	"gopkg.in/sqle/vitess-go.v2/event/syslogger"
)

// Syslog writes the event to syslog.
func (kc *KeyspaceChange) Syslog() (syslog.Priority, string) {
	return syslog.LOG_INFO, fmt.Sprintf("%s [keyspace] %s value: %s",
		kc.KeyspaceName, kc.Status, kc.Keyspace.String())
}

var _ syslogger.Syslogger = (*KeyspaceChange)(nil) // compile-time interface check
