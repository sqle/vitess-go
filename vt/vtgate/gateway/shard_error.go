// Copyright 2016, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gateway

import (
	"gopkg.in/sqle/vitess-go.v1/vt/topo/topoproto"
	"gopkg.in/sqle/vitess-go.v1/vt/vterrors"

	querypb "gopkg.in/sqle/vitess-go.v1/vt/proto/query"
	topodatapb "gopkg.in/sqle/vitess-go.v1/vt/proto/topodata"
)

// NewShardError returns a new error with the shard info amended.
func NewShardError(in error, target *querypb.Target, tablet *topodatapb.Tablet, inTransaction bool) error {
	if in == nil {
		return nil
	}
	if tablet != nil {
		return vterrors.Errorf(vterrors.Code(in), "target: %s.%s.%s, used tablet: (%+v), %v", target.Keyspace, target.Shard, topoproto.TabletTypeLString(target.TabletType), tablet, in)
	}
	return vterrors.Errorf(vterrors.Code(in), "target: %s.%s.%s, %v", target.Keyspace, target.Shard, topoproto.TabletTypeLString(target.TabletType), in)
}
