// Copyright 2013, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Imports and register the gRPC queryservice server

import (
	"gopkg.in/sqle/vitess-go.v1/vt/servenv"
	"gopkg.in/sqle/vitess-go.v1/vt/tabletserver"
	"gopkg.in/sqle/vitess-go.v1/vt/tabletserver/grpcqueryservice"
)

func init() {
	servenv.RegisterGRPCFlags()
	tabletserver.RegisterFunctions = append(tabletserver.RegisterFunctions, func(qsc tabletserver.Controller) {
		if servenv.GRPCCheckServiceMap("queryservice") {
			grpcqueryservice.Register(servenv.GRPCServer, qsc.QueryService())
		}
	})
}
