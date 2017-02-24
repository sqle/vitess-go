// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goclienttest

import (
	"encoding/json"
	"strings"
	"testing"

	"golang.org/x/net/context"
	"gopkg.in/sqle/vitess-go.v1/cmd/vtgateclienttest/services"
	"gopkg.in/sqle/vitess-go.v1/vt/callerid"
	"gopkg.in/sqle/vitess-go.v1/vt/vtgate/vtgateconn"

	querypb "gopkg.in/sqle/vitess-go.v1/vt/proto/query"
	topodatapb "gopkg.in/sqle/vitess-go.v1/vt/proto/topodata"
	vtgatepb "gopkg.in/sqle/vitess-go.v1/vt/proto/vtgate"
)

// testCallerID adds a caller ID to a context, and makes sure the server
// gets it.
func testCallerID(t *testing.T, conn *vtgateconn.VTGateConn) {
	t.Log("testCallerID")
	ctx := context.Background()
	callerID := callerid.NewEffectiveCallerID("test_principal", "test_component", "test_subcomponent")
	ctx = callerid.NewContext(ctx, callerID, nil)

	data, err := json.Marshal(callerID)
	if err != nil {
		t.Errorf("failed to marshal callerid: %v", err)
		return
	}
	query := services.CallerIDPrefix + string(data)

	// test Execute calls forward the callerID
	_, err = conn.Execute(ctx, query, nil, topodatapb.TabletType_MASTER, nil)
	checkCallerIDError(t, "Execute", err)

	_, err = conn.ExecuteShards(ctx, query, "", nil, nil, topodatapb.TabletType_MASTER, nil)
	checkCallerIDError(t, "ExecuteShards", err)

	_, err = conn.ExecuteKeyspaceIds(ctx, query, "", nil, nil, topodatapb.TabletType_MASTER, nil)
	checkCallerIDError(t, "ExecuteKeyspaceIds", err)

	_, err = conn.ExecuteKeyRanges(ctx, query, "", nil, nil, topodatapb.TabletType_MASTER, nil)
	checkCallerIDError(t, "ExecuteKeyRanges", err)

	_, err = conn.ExecuteEntityIds(ctx, query, "", "", nil, nil, topodatapb.TabletType_MASTER, nil)
	checkCallerIDError(t, "ExecuteEntityIds", err)

	// test ExecuteBatch calls forward the callerID
	_, err = conn.ExecuteBatchShards(ctx, []*vtgatepb.BoundShardQuery{
		{
			Query: &querypb.BoundQuery{
				Sql: query,
			},
		},
	}, topodatapb.TabletType_MASTER, false, nil)
	checkCallerIDError(t, "ExecuteBatchShards", err)

	_, err = conn.ExecuteBatchKeyspaceIds(ctx, []*vtgatepb.BoundKeyspaceIdQuery{
		{
			Query: &querypb.BoundQuery{
				Sql: query,
			},
		},
	}, topodatapb.TabletType_MASTER, false, nil)
	checkCallerIDError(t, "ExecuteBatchKeyspaceIds", err)

	// test StreamExecute calls forward the callerID
	err = getStreamError(conn.StreamExecute(ctx, query, nil, topodatapb.TabletType_MASTER, nil))
	checkCallerIDError(t, "StreamExecute", err)

	err = getStreamError(conn.StreamExecuteShards(ctx, query, "", nil, nil, topodatapb.TabletType_MASTER, nil))
	checkCallerIDError(t, "StreamExecuteShards", err)

	err = getStreamError(conn.StreamExecuteKeyspaceIds(ctx, query, "", nil, nil, topodatapb.TabletType_MASTER, nil))
	checkCallerIDError(t, "StreamExecuteKeyspaceIds", err)

	err = getStreamError(conn.StreamExecuteKeyRanges(ctx, query, "", nil, nil, topodatapb.TabletType_MASTER, nil))
	checkCallerIDError(t, "StreamExecuteKeyRanges", err)

	// test UpdateStream forwards the callerID
	err = getUpdateStreamError(conn.UpdateStream(ctx, query, nil, topodatapb.TabletType_MASTER, 0, nil))
	checkCallerIDError(t, "UpdateStream", err)
}

func checkCallerIDError(t *testing.T, name string, err error) {
	if err == nil {
		t.Errorf("callerid: got no error for %v", name)
		return
	}
	if !strings.Contains(err.Error(), "SUCCESS: ") {
		t.Errorf("failed to pass callerid for %v: %v", name, err)
	}
}
