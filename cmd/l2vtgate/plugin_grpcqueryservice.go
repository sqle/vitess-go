// Copyright 2013, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Imports and register the gRPC queryservice server

import (
	"gopkg.in/sqle/vitess-go.v2/vt/servenv"
	"gopkg.in/sqle/vitess-go.v2/vt/vtgate/l2vtgate"
	"gopkg.in/sqle/vitess-go.v2/vt/vttablet/grpcqueryservice"
	"gopkg.in/sqle/vitess-go.v2/vt/vttablet/queryservice"
)

func init() {
	servenv.RegisterGRPCFlags()
	l2vtgate.RegisterL2VTGates = append(l2vtgate.RegisterL2VTGates, func(qs queryservice.QueryService) {
		if servenv.GRPCCheckServiceMap("queryservice") {
			grpcqueryservice.Register(servenv.GRPCServer, qs)
		}
	})
}
