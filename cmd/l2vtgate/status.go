package main

import (
	"gopkg.in/sqle/vitess-go.v1/vt/discovery"
	"gopkg.in/sqle/vitess-go.v1/vt/servenv"
	_ "gopkg.in/sqle/vitess-go.v1/vt/status"
	"gopkg.in/sqle/vitess-go.v1/vt/vtgate"
	"gopkg.in/sqle/vitess-go.v1/vt/vtgate/gateway"
	"gopkg.in/sqle/vitess-go.v1/vt/vtgate/l2vtgate"
)

// For use by plugins which wish to avoid racing when registering status page parts.
var onStatusRegistered func()

func addStatusParts(l2vtgate *l2vtgate.L2VTGate) {
	servenv.AddStatusPart("Topology Cache", vtgate.TopoTemplate, func() interface{} {
		return resilientSrvTopoServer.CacheStatus()
	})
	servenv.AddStatusPart("Gateway Status", gateway.StatusTemplate, func() interface{} {
		return l2vtgate.GetGatewayCacheStatus()
	})
	servenv.AddStatusPart("Health Check Cache", discovery.HealthCheckTemplate, func() interface{} {
		return healthCheck.CacheStatus()
	})
	if onStatusRegistered != nil {
		onStatusRegistered()
	}
}
