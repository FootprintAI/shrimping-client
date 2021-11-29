package version

import "fmt"

var (
	BuildTime   = ""
	GitCommitId = ""
	version     = "1.0.3"
)

func GetVersion() string {
	return version
}

func GetBuildTime() string {
	return BuildTime
}

func GetCommitHash() string {
	return GitCommitId
}

func Print() {
	fmt.Printf("========shrimping cli=======\n")
	fmt.Printf("commit hash: %s\nbuild time: %s\nclient verison: %s\n", GitCommitId, BuildTime, version)
	fmt.Printf("============================\n")
}
