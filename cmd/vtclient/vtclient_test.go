package main

import (
	"flag"
	"os"
	"strings"
	"testing"

	"gopkg.in/sqle/vitess-go.v2/vt/vttest"

	vschemapb "gopkg.in/sqle/vitess-go.v2/vt/proto/vschema"
	vttestpb "gopkg.in/sqle/vitess-go.v2/vt/proto/vttest"
)

func TestVtclient(t *testing.T) {
	topology := &vttestpb.VTTestTopology{
		Keyspaces: []*vttestpb.Keyspace{
			{
				Name: "test_keyspace",
				Shards: []*vttestpb.Shard{
					{
						Name: "0",
					},
				},
			},
		},
	}
	schema := `CREATE TABLE table1 (
	  id BIGINT(20) UNSIGNED NOT NULL,
	  i INT NOT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB`
	vschema := &vschemapb.Keyspace{
		Sharded: true,
		Vindexes: map[string]*vschemapb.Vindex{
			"hash": {
				Type: "hash",
			},
		},
		Tables: map[string]*vschemapb.Table{
			"table1": {
				ColumnVindexes: []*vschemapb.ColumnVindex{
					{
						Column: "id",
						Name:   "hash",
					},
				},
			},
		},
	}

	hdl, err := vttest.LaunchVitess(vttest.ProtoTopo(topology), vttest.Schema(schema), vttest.VSchema(vschema))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = hdl.TearDown()
		if err != nil {
			t.Fatal(err)
		}
	}()

	vtgateAddr, err := hdl.VtgateAddress()
	if err != nil {
		t.Fatal(err)
	}

	queries := []struct {
		args         []string
		rowsAffected int64
		errMsg       string
	}{
		{
			args: []string{"SELECT * FROM table1"},
		},
		{
			args: []string{"-tablet_type", "master", "-bind_variables", `[ 1, 100 ]`,
				"INSERT INTO table1 (id, i) VALUES (:v1, :v2)"},
			rowsAffected: 1,
		},
		{
			args: []string{"-tablet_type", "master",
				"UPDATE table1 SET i = (i + 1)"},
			rowsAffected: 1,
		},
		{
			args: []string{"-tablet_type", "master",
				"SELECT * FROM table1"},
			rowsAffected: 1,
		},
		{
			args: []string{"-tablet_type", "master", "-bind_variables", `[ 1 ]`,
				"DELETE FROM table1 WHERE id = :v1"},
			rowsAffected: 1,
		},
		{
			args: []string{"-tablet_type", "master",
				"SELECT * FROM table1"},
			rowsAffected: 0,
		},
		{
			args:   []string{"SELECT * FROM nonexistent"},
			errMsg: "table nonexistent not found in schema",
		},
	}

	// Change ErrorHandling from ExitOnError to panicking.
	flag.CommandLine.Init("vtclient_test.go", flag.PanicOnError)
	for _, q := range queries {
		// Run main function directly and not as external process. To achieve this,
		// overwrite os.Args which is used by flag.Parse().
		os.Args = []string{"vtclient_test.go", "-server", vtgateAddr}
		os.Args = append(os.Args, q.args...)

		results, err := run()
		if q.errMsg != "" {
			if got, want := err.Error(), q.errMsg; !strings.Contains(got, want) {
				t.Fatalf("vtclient %v returned wrong error: got = %v, want contains = %v", os.Args[1:], got, want)
			}
			return
		}

		if err != nil {
			t.Fatalf("vtclient %v failed: %v", os.Args[1:], err)
		}
		if got, want := results.rowsAffected, q.rowsAffected; got != want {
			t.Fatalf("wrong rows affected for query: %v got = %v, want = %v", os.Args[1:], got, want)
		}
	}
}
