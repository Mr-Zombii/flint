package version

import (
	"fmt"
	"runtime"
)

var (
	Version   = "unknown"
	GitCommit = "unknown"
	BuildTime = "unknown"
	GoVersion = runtime.Version()
	OS        = runtime.GOOS
	Arch      = runtime.GOARCH
)

func FullVersion() string {
	return fmt.Sprintf(
		"Flint Version: %s\nCommit: %s\nBuild Time: %s\nGo Version: %s\nPlatform: %s/%s",
		Version, GitCommit, BuildTime, GoVersion, OS, Arch,
	)
}
