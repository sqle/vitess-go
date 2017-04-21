package events

import (
	topodatapb "gopkg.in/sqle/vitess-go.v2/vt/proto/topodata"
)

// KeyspaceChange is an event that describes changes to a keyspace.
type KeyspaceChange struct {
	KeyspaceName string
	Keyspace     *topodatapb.Keyspace
	Status       string
}
