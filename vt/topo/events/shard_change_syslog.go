package events

import (
	"fmt"
	"log/syslog"

	"gopkg.in/sqle/vitess-go.v2/event/syslogger"
)

// Syslog writes the event to syslog.
func (sc *ShardChange) Syslog() (syslog.Priority, string) {
	return syslog.LOG_INFO, fmt.Sprintf("%s/%s [shard] %s value: %s",
		sc.KeyspaceName, sc.ShardName, sc.Status, sc.Shard.String())
}

var _ syslogger.Syslogger = (*ShardChange)(nil) // compile-time interface check
