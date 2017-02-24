// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package services exposes all the services for the vtgateclienttest binary.
package services

import "gopkg.in/sqle/vitess-go.v1/vt/vtgate/vtgateservice"

// CreateServices creates the implementation chain of all the test cases
func CreateServices() vtgateservice.VTGateService {
	var s vtgateservice.VTGateService
	s = newTerminalClient()
	s = newSuccessClient(s)
	s = newErrorClient(s)
	s = newCallerIDClient(s)
	s = newEchoClient(s)
	return s
}
