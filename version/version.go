package version

import (
	_ "embed"
	"fmt"
	"runtime/debug"
)

// GetVersion returns the Version for the current binary.
func GetVersion() string {

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "UNKNOWN (no ReadBuildInfo())"
	}

	settings := map[string]string{}
	for _, setting := range info.Settings {
		settings[setting.Key] = setting.Value
	}

	revision, ok := settings["vcs.revision"]
	if !ok {
		return `UNKNOWN (no "vcs.revision")`
	}

	version := revision

	modified, ok := settings["vcs.modified"]
	if ok && modified == "true" {
		version = fmt.Sprintf("%s (dirty)", version)
	}

	return version
}
