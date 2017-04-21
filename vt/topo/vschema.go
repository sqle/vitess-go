package topo

import (
	"golang.org/x/net/context"

	vschemapb "gopkg.in/sqle/vitess-go.v2/vt/proto/vschema"
	"gopkg.in/sqle/vitess-go.v2/vt/vtgate/vindexes"
)

// SaveVSchema first validates the VSchema, then sends it to the underlying
// Impl.
func (ts Server) SaveVSchema(ctx context.Context, keyspace string, vschema *vschemapb.Keyspace) error {
	err := vindexes.ValidateKeyspace(vschema)
	if err != nil {
		return err
	}

	return ts.Impl.SaveVSchema(ctx, keyspace, vschema)
}
