package pkg

import "fmt"

// Version component constants for the current build.
const (
	VersionMajor         = 0
	VersionMinor         = 20
	VersionPatch         = 0
	VersionReleaseLevel  = "alpha"
	VersionReleaseNumber = 23
)

// Set the GitVersion via -ldflags="-X 'github.com/trisacrypto/envoy/pkg.GitVersion=$(git rev-parse --short HEAD)'"
var GitVersion string

// Version returns the semantic version for the current build.
func Version() string {
	versionCore := fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)

	if VersionReleaseLevel != "" {
		if VersionReleaseNumber > 0 {
			versionCore = fmt.Sprintf("%s-%s.%d", versionCore, VersionReleaseLevel, VersionReleaseNumber)
		} else {
			versionCore = fmt.Sprintf("%s-%s", versionCore, VersionReleaseLevel)
		}
	}

	if GitVersion != "" {
		versionCore = fmt.Sprintf("%s (%s)", versionCore, GitVersion)
	}

	return versionCore
}
