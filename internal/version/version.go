package version

import (
	"fmt"
	"runtime/debug"
)

// Set at link time (see Makefile / Dockerfile).
var (
	Version   = "dev"
	Commit    = ""
	BuildDate = ""
)

// String returns a single-line build description for `-version` and logs.
// When Commit is empty, it tries [debug.ReadBuildInfo] (e.g. go install from VCS).
func String() string {
	rev := Commit
	if rev == "" {
		rev = vcsRevisionFromBuildInfo()
	}
	if rev == "" {
		rev = "unknown"
	} else if len(rev) > 7 {
		rev = rev[:7]
	}

	date := BuildDate
	if date == "" {
		date = "unknown"
	}

	return fmt.Sprintf("%s (commit %s, %s)", Version, rev, date)
}

func vcsRevisionFromBuildInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			return s.Value
		}
	}
	return ""
}
