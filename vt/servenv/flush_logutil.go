package servenv

import (
	"gopkg.in/sqle/vitess-go.v2/vt/logutil"
)

func init() {
	OnClose(logutil.Flush)
}
