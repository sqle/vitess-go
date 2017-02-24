// Copyright 2017, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabletconn

import (
	"testing"

	vtrpcpb "gopkg.in/sqle/vitess-go.v1/vt/proto/vtrpc"
	"gopkg.in/sqle/vitess-go.v1/vt/vterrors"
)

func TestTabletErrorFromRPCError(t *testing.T) {
	testcases := []struct {
		in   *vtrpcpb.RPCError
		want vtrpcpb.Code
	}{{
		in: &vtrpcpb.RPCError{
			LegacyCode: vtrpcpb.LegacyErrorCode_BAD_INPUT_LEGACY,
			Message:    "bad input",
		},
		want: vtrpcpb.Code_INVALID_ARGUMENT,
	}, {
		in: &vtrpcpb.RPCError{
			LegacyCode: vtrpcpb.LegacyErrorCode_BAD_INPUT_LEGACY,
			Message:    "bad input",
			Code:       vtrpcpb.Code_INVALID_ARGUMENT,
		},
		want: vtrpcpb.Code_INVALID_ARGUMENT,
	}, {
		in: &vtrpcpb.RPCError{
			Message: "bad input",
			Code:    vtrpcpb.Code_INVALID_ARGUMENT,
		},
		want: vtrpcpb.Code_INVALID_ARGUMENT,
	}}
	for _, tcase := range testcases {
		got := vterrors.Code(ErrorFromVTRPC(tcase.in))
		if got != tcase.want {
			t.Errorf("FromVtRPCError(%v):\n%v, want\n%v", tcase.in, got, tcase.want)
		}
	}
}
