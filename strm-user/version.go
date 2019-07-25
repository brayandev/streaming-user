package user

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

// filled during build time using linking
var (
	AppVersion  = ""
	BuildHash   = ""
	BuildNumber = ""
	BuildDate   = ""
)

// startTime application
var startTime = time.Now().Format(time.RFC3339)

// BuildPlataform OS/ARCH
var BuildPlataform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

// VersionString format version as string
func VersionString() string {
	return fmt.Sprintf("Os: %s\n AppVersion: %s\nBuildHash: %s\nBuildNumber: %s\nBuildDate: %s\nBuildPlataform: %s\nGoVersion: %s",
		runtime.GOOS, AppVersion, BuildHash, BuildNumber, BuildDate, BuildPlataform, runtime.Version())
}

// JSON format version as json.
func JSON() []byte {
	content := map[string]string{
		"os":          runtime.GOOS,
		"arch":        runtime.GOARCH,
		"startTime":   startTime,
		"appVersion":  AppVersion,
		"goVersion":   runtime.Version(),
		"buildHash":   BuildHash,
		"buildNumber": BuildNumber,
		"buildDate":   BuildDate,
	}

	result, _ := json.Marshal(content)
	return result
}
