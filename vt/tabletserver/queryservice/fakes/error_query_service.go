package fakes

import (
	"fmt"

	"golang.org/x/net/context"
	"gopkg.in/sqle/vitess-go.v1/vt/tabletserver/queryservice"

	querypb "gopkg.in/sqle/vitess-go.v1/vt/proto/query"
)

// ErrorQueryService is an object that returns an error for all methods.
var ErrorQueryService = queryservice.Wrap(
	nil,
	func(ctx context.Context, target *querypb.Target, conn queryservice.QueryService, name string, inTransaction bool, inner func(context.Context, *querypb.Target, queryservice.QueryService) (error, bool)) error {
		return fmt.Errorf("ErrorQueryService does not implement any method")
	},
)
