package build

import "runtime"

//nolint:gochecknoglobals // these variables is set during the build
var (
	Version       = "unknown"
	VersionCommit = "unknown"
	BuildDate     = "unknown"
	GoVersion     = runtime.Version()
)
