// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"os"
	"time"

	log "github.com/golang/glog"
	"golang.org/x/net/context"
	"gopkg.in/sqle/vitess-go.v1/exit"
	"gopkg.in/sqle/vitess-go.v1/vt/logutil"
	"gopkg.in/sqle/vitess-go.v1/vt/vtctl/vtctlclient"

	logutilpb "gopkg.in/sqle/vitess-go.v1/vt/proto/logutil"
)

// The default values used by these flags cannot be taken from wrangler and
// actionnode modules, as we do't want to depend on them at all.
var (
	actionTimeout = flag.Duration("action_timeout", time.Hour, "timeout for the total command")
	dialTimeout   = flag.Duration("dial_timeout", 30*time.Second, "time to wait for the dial phase")
	server        = flag.String("server", "", "server to use for connection")
)

func main() {
	defer exit.Recover()

	flag.Parse()

	logger := logutil.NewConsoleLogger()

	err := vtctlclient.RunCommandAndWait(
		context.Background(), *server, flag.Args(),
		*dialTimeout, *actionTimeout,
		func(e *logutilpb.Event) {
			logutil.LogEvent(logger, e)
		})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
