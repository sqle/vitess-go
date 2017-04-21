// Copyright 2014, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package events defines the structures used for events dispatched from the
// wrangler package.
package events

import (
	base "gopkg.in/sqle/vitess-go.v2/vt/events"
	"gopkg.in/sqle/vitess-go.v2/vt/topo"

	topodatapb "gopkg.in/sqle/vitess-go.v2/vt/proto/topodata"
)

// Reparent is an event that describes a single step in the reparent process.
type Reparent struct {
	base.StatusUpdater

	ShardInfo            topo.ShardInfo
	OldMaster, NewMaster topodatapb.Tablet
	ExternalID           string
}
