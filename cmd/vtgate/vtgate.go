// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"

	"gopkg.in/sqle/vitess-go.v1/exit"
	"gopkg.in/sqle/vitess-go.v1/vt/discovery"
	"gopkg.in/sqle/vitess-go.v1/vt/servenv"
	"gopkg.in/sqle/vitess-go.v1/vt/topo"
	"gopkg.in/sqle/vitess-go.v1/vt/topo/topoproto"
	"gopkg.in/sqle/vitess-go.v1/vt/vtgate"

	topodatapb "gopkg.in/sqle/vitess-go.v1/vt/proto/topodata"
)

var (
	cell                   = flag.String("cell", "test_nj", "cell to use")
	retryCount             = flag.Int("retry-count", 2, "retry count")
	healthCheckConnTimeout = flag.Duration("healthcheck_conn_timeout", 3*time.Second, "healthcheck connection timeout")
	healthCheckRetryDelay  = flag.Duration("healthcheck_retry_delay", 2*time.Millisecond, "health check retry delay")
	healthCheckTimeout     = flag.Duration("healthcheck_timeout", time.Minute, "the health check timeout period")
	tabletTypesToWait      = flag.String("tablet_types_to_wait", "", "wait till connected for specified tablet types during Gateway initialization")
)

var resilientSrvTopoServer *vtgate.ResilientSrvTopoServer
var healthCheck discovery.HealthCheck

var initFakeZK func()

func init() {
	rand.Seed(time.Now().UnixNano())
	servenv.RegisterDefaultFlags()
}

func main() {
	defer exit.Recover()

	flag.Parse()
	servenv.Init()

	if initFakeZK != nil {
		initFakeZK()
	}
	ts := topo.Open()
	defer ts.Close()

	resilientSrvTopoServer = vtgate.NewResilientSrvTopoServer(ts, "ResilientSrvTopoServer")

	healthCheck = discovery.NewHealthCheck(*healthCheckConnTimeout, *healthCheckRetryDelay, *healthCheckTimeout)
	healthCheck.RegisterStats()

	tabletTypes := make([]topodatapb.TabletType, 0, 1)
	if len(*tabletTypesToWait) != 0 {
		for _, ttStr := range strings.Split(*tabletTypesToWait, ",") {
			tt, err := topoproto.ParseTabletType(ttStr)
			if err != nil {
				log.Errorf("unknown tablet type: %v", ttStr)
				continue
			}
			tabletTypes = append(tabletTypes, tt)
		}
	}

	vtg := vtgate.Init(context.Background(), healthCheck, ts, resilientSrvTopoServer, *cell, *retryCount, tabletTypes)

	servenv.OnRun(func() {
		addStatusParts(vtg)
	})
	servenv.RunDefault()
}
