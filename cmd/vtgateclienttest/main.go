// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main is the implementation of vtgateclienttest.
// This program has a chain of vtgateservice.VTGateService implementations,
// each one being responsible for one test scenario.
package main

import (
	"flag"

	"gopkg.in/sqle/vitess-go.v1/cmd/vtgateclienttest/services"
	"gopkg.in/sqle/vitess-go.v1/exit"
	"gopkg.in/sqle/vitess-go.v1/vt/servenv"
	"gopkg.in/sqle/vitess-go.v1/vt/vtgate"
)

func init() {
	servenv.RegisterDefaultFlags()
}

func main() {
	defer exit.Recover()

	flag.Parse()
	servenv.Init()

	// The implementation chain.
	servenv.OnRun(func() {
		s := services.CreateServices()
		for _, f := range vtgate.RegisterVTGates {
			f(s)
		}
	})

	servenv.RunDefault()
}
