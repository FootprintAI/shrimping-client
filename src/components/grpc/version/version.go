package version

import (
	"fmt"

	goversion "github.com/hashicorp/go-version"
)

var (
	BuildTime   = ""
	GitCommitId = ""
	version, _  = goversion.NewVersion("1.2.0")
)

func GreatThan(v1, v2 string) bool {
	v1v, _ := goversion.NewVersion(v1)
	v2v, _ := goversion.NewVersion(v2)

	return v1v.GreaterThan(v2v)
}

func GetVersion() string {
	return version.String()
}

func GetBuildTime() string {
	return BuildTime
}

func GetCommitHash() string {
	return GitCommitId
}

func Print() {
	fmt.Printf("ver:%s, build time:%s, hashid:%s\n", version.String(), BuildTime, GitCommitId)
}
