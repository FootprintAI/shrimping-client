package version

import (
	"fmt"
	"strconv"
	"strings"
)

func Great(s1, s2 string) bool {
	s1v := convertVer2Numeric(s1)
	s2v := convertVer2Numeric(s2)
	return s1v > s2v
}

func convertVer2Numeric(s string) int {
	parts := strings.SplitN(s, ".", 3)
	return mustConvert2Int(parts[0])*1000*1000 + mustConvert2Int(parts[1])*1000 + mustConvert2Int(parts[2])
}

func mustConvert2Int(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

var (
	BuildTime   = ""
	GitCommitId = ""
	version     = "1.1.5"
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
