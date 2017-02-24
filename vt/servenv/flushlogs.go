package servenv

import (
	"fmt"
	"net/http"

	"gopkg.in/sqle/vitess-go.v1/vt/logutil"
)

func init() {
	onInit(func() {
		http.HandleFunc("/debug/flushlogs", func(w http.ResponseWriter, r *http.Request) {
			logutil.Flush()
			fmt.Fprint(w, "flushed")
		})
	})
}
