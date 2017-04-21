package fakes

import (
	"fmt"

	"gopkg.in/sqle/vitess-go.v2/vt/vttablet/queryservice"
	"golang.org/x/net/context"

	querypb "gopkg.in/sqle/vitess-go.v2/vt/proto/query"
)

// ErrorQueryService is an object that returns an error for all methods.
var ErrorQueryService = queryservice.Wrap(
	nil,
	func(ctx context.Context, target *querypb.Target, conn queryservice.QueryService, name string, inTransaction bool, inner func(context.Context, *querypb.Target, queryservice.QueryService) (error, bool)) error {
		return fmt.Errorf("ErrorQueryService does not implement any method")
	},
)
