// Package eventtoken includes utility methods for event token
// handling.
package eventtoken

import (
	"gopkg.in/sqle/vitess-go.v2/mysqlconn/replication"

	querypb "gopkg.in/sqle/vitess-go.v2/vt/proto/query"
)

// Fresher compares two event tokens.  It returns a negative number if
// ev1<ev2, zero if they're equal, and a positive number if
// ev1>ev2. In case of doubt (we don't have enough information to know
// for sure), it returns a negative number.
func Fresher(ev1, ev2 *querypb.EventToken) int {
	if ev1 == nil || ev2 == nil {
		// Either one is nil, we don't know.
		return -1
	}

	if ev1.Timestamp != ev2.Timestamp {
		// The timestamp is enough to set them apart.
		return int(ev1.Timestamp - ev2.Timestamp)
	}

	if ev1.Shard != "" && ev1.Shard == ev2.Shard {
		// They come from the same shard. See if we have positions.
		if ev1.Position == "" || ev2.Position == "" {
			return -1
		}

		// We can parse them.
		pos1, err := replication.DecodePosition(ev1.Position)
		if err != nil {
			return -1
		}
		pos2, err := replication.DecodePosition(ev2.Position)
		if err != nil {
			return -1
		}

		// Then compare.
		if pos1.Equal(pos2) {
			return 0
		}
		if pos1.AtLeast(pos2) {
			return 1
		}
		return -1
	}

	// We do not know.
	return -1
}
