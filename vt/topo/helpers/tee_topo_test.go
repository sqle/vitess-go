// Copyright 2013, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package helpers

import (
	"testing"

	"gopkg.in/sqle/vitess-go.v2/vt/topo"
	"gopkg.in/sqle/vitess-go.v2/vt/topo/memorytopo"
	"gopkg.in/sqle/vitess-go.v2/vt/topo/test"
)

func newFakeTeeServer(t *testing.T) topo.Impl {
	s1 := memorytopo.New("test")
	s2 := memorytopo.New("test")
	return NewTee(s1, s2, false)
}

func TestTeeTopo(t *testing.T) {
	test.TopoServerTestSuite(t, func() topo.Impl {
		return newFakeTeeServer(t)
	})
}
