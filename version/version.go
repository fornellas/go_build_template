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
		return "unknown revision (no ReadBuildInfo())"
	}

	settings := map[string]string{}
	for _, setting := range info.Settings {
		settings[setting.Key] = setting.Value
	}

	revision, ok := settings["vcs.revision"]
	if !ok {
		return `unknown revision (no "vcs.revision")`
	}

	modified, ok := settings["vcs.modified"]
	if ok && modified == "true" {
		revision = fmt.Sprintf("%s (modified)", revision)
	}

	version := fmt.Sprintf("%s (%s)", info.Main.Version, revision)

	return version
}
