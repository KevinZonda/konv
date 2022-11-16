package path

import (
	"os"
	"path"
	"runtime"
)

func getBase() string {
	if runtime.GOOS == "windows" {
		dir, _ := os.UserConfigDir()
		return path.Join(dir, "konv")
	}
	return "/etc/konv"
}

func getSelfBase() string {
	dir, _ := os.UserHomeDir()
	return path.Join(dir, ".config", "konv")
}

func GetConvPath(target string, isSelf bool) string {
	var base string
	if isSelf {
		base = getSelfBase()
	} else {
		base = getBase()
	}
	return path.Join(base, target+".csv")
}

func GetConvCfgPath(target string, isSelf bool) string {
	var base string
	if isSelf {
		base = getSelfBase()
	} else {
		base = getBase()
	}
	return path.Join(base, target+".cfg")
}
