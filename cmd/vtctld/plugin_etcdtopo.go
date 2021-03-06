// Copyright 2014, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Imports and register the 'etcd' topo.Server and its Explorer.

import (
	"gopkg.in/sqle/vitess-go.v2/vt/etcdtopo"
	"gopkg.in/sqle/vitess-go.v2/vt/servenv"
	"gopkg.in/sqle/vitess-go.v2/vt/vtctld"
)

func init() {
	// Wait until flags are parsed, so we can check which topo server is in use.
	servenv.OnRun(func() {
		if etcdServer, ok := ts.Impl.(*etcdtopo.Server); ok {
			vtctld.HandleExplorer("etcd", etcdtopo.NewExplorer(etcdServer))
		}
	})
}
