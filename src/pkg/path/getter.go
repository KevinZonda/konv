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

func GetConversionPath(target string) string {
	return path.Join(getBase(), target+".csv")
}

func GetConversionCfg(target string) string {
	return path.Join(getBase(), target+".cfg")
}
