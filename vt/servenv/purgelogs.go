package servenv

import (
	"gopkg.in/sqle/vitess-go.v1/vt/logutil"
)

func init() {
	onInit(func() {
		go logutil.PurgeLogs()
	})

}
