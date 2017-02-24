package schema

import (
	"fmt"
	"testing"

	querypb "gopkg.in/sqle/vitess-go.v1/vt/proto/query"
	"gopkg.in/sqle/vitess-go.v1/vt/sqlparser"
)

func TestTableColumnString(t *testing.T) {
	c := &TableColumn{Name: sqlparser.NewColIdent("my_column"), Type: querypb.Type_INT8}
	want := "{Name: 'my_column', Type: INT8}"
	got := fmt.Sprintf("%v", c)
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
