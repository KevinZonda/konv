package path

import (
	"os"
	"path"
	"runtime"
)

func GetConversionPath(target string) string {
	if runtime.GOOS == "windows" {
		dir, _ := os.UserConfigDir()
		return path.Join(dir, "konv", target+".csv")
	}
	return "/etc/konv/" + target + ".csv"
}
