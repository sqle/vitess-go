package servenv

import (
	"fmt"
	"time"

	"gopkg.in/sqle/vitess-go.v1/stats"
)

var (
	buildHost   = ""
	buildUser   = ""
	buildTime   = ""
	buildGitRev = ""
)

func init() {
	t, err := time.Parse(time.UnixDate, buildTime)
	if buildTime != "" && err != nil {
		panic(fmt.Sprintf("Couldn't parse build timestamp %q: %v", buildTime, err))
	}
	stats.NewString("BuildHost").Set(buildHost)
	stats.NewString("BuildUser").Set(buildUser)
	stats.NewInt("BuildTimestamp").Set(t.Unix())
	stats.NewString("BuildGitRev").Set(buildGitRev)
}
