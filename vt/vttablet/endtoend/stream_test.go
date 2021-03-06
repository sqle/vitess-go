// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package endtoend

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"gopkg.in/sqle/vitess-go.v2/sqltypes"
	vtrpcpb "gopkg.in/sqle/vitess-go.v2/vt/proto/vtrpc"
	"gopkg.in/sqle/vitess-go.v2/vt/vttablet/endtoend/framework"
	"gopkg.in/sqle/vitess-go.v2/vt/vterrors"
)

func TestStreamUnion(t *testing.T) {
	qr, err := framework.NewClient().StreamExecute("select 1 from dual union select 1 from dual", nil)
	if err != nil {
		t.Error(err)
		return
	}
	if qr.RowsAffected != 1 {
		t.Errorf("RowsAffected: %d, want 1", qr.RowsAffected)
	}
}

func TestStreamBigData(t *testing.T) {
	client := framework.NewClient()
	err := populateBigData(client)
	if err != nil {
		t.Error(err)
		return
	}
	defer client.Execute("delete from vitess_big", nil)

	qr, err := client.StreamExecute("select * from vitess_big b1, vitess_big b2 order by b1.id, b2.id", nil)
	if err != nil {
		t.Error(err)
		return
	}
	row10 := framework.RowsToStrings(qr)[10]
	want := []string{
		"0",
		"AAAAAAAAAAAAAAAAAA 0",
		"BBBBBBBBBBBBBBBBBB 0",
		"C",
		"DDDDDDDDDDDDDDDDDD 0",
		"EEEEEEEEEEEEEEEEEE 0",
		"FF 0",
		"GGGGGGGGGGGGGGGGGG 0",
		"0",
		"0",
		"0",
		"0",
		"10",
		"AAAAAAAAAAAAAAAAAA 10",
		"BBBBBBBBBBBBBBBBBB 10",
		"C",
		"DDDDDDDDDDDDDDDDDD 10",
		"EEEEEEEEEEEEEEEEEE 10",
		"FF 10",
		"GGGGGGGGGGGGGGGGGG 10",
		"10",
		"10",
		"10",
		"10"}
	if !reflect.DeepEqual(row10, want) {
		t.Errorf("Row10: \n%#v, want \n%#v", row10, want)
	}
}

func TestStreamTerminate(t *testing.T) {
	client := framework.NewClient()
	err := populateBigData(client)
	if err != nil {
		t.Error(err)
		return
	}
	defer client.Execute("delete from vitess_big", nil)

	called := false
	err = client.Stream(
		"select * from vitess_big b1, vitess_big b2 order by b1.id, b2.id",
		nil,
		func(*sqltypes.Result) error {
			if !called {
				queries := framework.StreamQueryz()
				if l := len(queries); l != 1 {
					t.Errorf("len(queries): %d, want 1", l)
					return errors.New("no queries from StreamQueryz")
				}
				err := framework.StreamTerminate(queries[0].ConnID)
				if err != nil {
					return err
				}
				called = true
			}
			time.Sleep(10 * time.Millisecond)
			return nil
		},
	)
	if code := vterrors.Code(err); code != vtrpcpb.Code_DEADLINE_EXCEEDED {
		t.Errorf("Errorcode: %v, want %v", code, vtrpcpb.Code_DEADLINE_EXCEEDED)
	}
}

func populateBigData(client *framework.QueryClient) error {
	err := client.Begin()
	if err != nil {
		return err
	}
	defer client.Rollback()

	for i := 0; i < 100; i++ {
		stri := strconv.Itoa(i)
		query := "insert into vitess_big values " +
			"(" + stri + ", " +
			"'AAAAAAAAAAAAAAAAAA " + stri + "', " +
			"'BBBBBBBBBBBBBBBBBB " + stri + "', " +
			"'C', " +
			"'DDDDDDDDDDDDDDDDDD " + stri + "', " +
			"'EEEEEEEEEEEEEEEEEE " + stri + "', " +
			"'FF " + stri + "', " +
			"'GGGGGGGGGGGGGGGGGG " + stri + "', " +
			stri + ", " +
			stri + ", " +
			stri + ", " +
			stri + ")"
		_, err := client.Execute(query, nil)
		if err != nil {
			return err
		}
	}
	return client.Commit()
}

func TestStreamError(t *testing.T) {
	_, err := framework.NewClient().StreamExecute("select count(abcd) from vitess_big", nil)
	want := "Unknown column"
	if err == nil || !strings.HasPrefix(err.Error(), want) {
		t.Errorf("Error: %v, must start with %s", err, want)
	}
}
