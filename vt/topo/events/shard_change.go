package events

import (
	topodatapb "gopkg.in/sqle/vitess-go.v1/vt/proto/topodata"
)

// ShardChange is an event that describes changes to a shard.
type ShardChange struct {
	KeyspaceName string
	ShardName    string
	Shard        *topodatapb.Shard
	Status       string
}
