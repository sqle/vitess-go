// Copyright 2014, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package grpcvtctlserver contains the gRPC implementation of the server side
of the remote execution of vtctl commands.
*/
package grpcvtctlserver

import (
	"google.golang.org/grpc"

	"gopkg.in/sqle/vitess-go.v1/vt/logutil"
	"gopkg.in/sqle/vitess-go.v1/vt/servenv"
	"gopkg.in/sqle/vitess-go.v1/vt/tabletmanager/tmclient"
	"gopkg.in/sqle/vitess-go.v1/vt/topo"
	"gopkg.in/sqle/vitess-go.v1/vt/vtctl"
	"gopkg.in/sqle/vitess-go.v1/vt/wrangler"

	logutilpb "gopkg.in/sqle/vitess-go.v1/vt/proto/logutil"
	vtctldatapb "gopkg.in/sqle/vitess-go.v1/vt/proto/vtctldata"
	vtctlservicepb "gopkg.in/sqle/vitess-go.v1/vt/proto/vtctlservice"
)

// VtctlServer is our RPC server
type VtctlServer struct {
	ts topo.Server
}

// NewVtctlServer returns a new Vtctl Server for the topo server.
func NewVtctlServer(ts topo.Server) *VtctlServer {
	return &VtctlServer{ts}
}

// ExecuteVtctlCommand is part of the vtctldatapb.VtctlServer interface
func (s *VtctlServer) ExecuteVtctlCommand(args *vtctldatapb.ExecuteVtctlCommandRequest, stream vtctlservicepb.Vtctl_ExecuteVtctlCommandServer) (err error) {
	defer servenv.HandlePanic("vtctl", &err)

	// create a logger, send the result back to the caller
	logstream := logutil.NewCallbackLogger(func(e *logutilpb.Event) {
		// If the client disconnects, we will just fail
		// to send the log events, but won't interrupt
		// the command.
		stream.Send(&vtctldatapb.ExecuteVtctlCommandResponse{
			Event: e,
		})
	})
	logger := logutil.NewTeeLogger(logstream, logutil.NewConsoleLogger())

	// create the wrangler
	tmc := tmclient.NewTabletManagerClient()
	defer tmc.Close()
	wr := wrangler.New(logger, s.ts, tmc)

	// execute the command
	return vtctl.RunCommand(stream.Context(), wr, args.Args)
}

// StartServer registers the VtctlServer for RPCs
func StartServer(s *grpc.Server, ts topo.Server) {
	vtctlservicepb.RegisterVtctlServer(s, NewVtctlServer(ts))
}
