package path

import (
	"os"
	"path"
	"runtime"
)

// base get base path
// Windows: %AppData%\konv
// Linux: /etc/konv
// MacOS: /etc/konv
func base() string {
	if runtime.GOOS == "windows" {
		dir, _ := os.UserConfigDir()
		return path.Join(dir, "konv")
	}
	return "/etc/konv"
}

// selfBase get self base path
// Windows: %USERPROFILE%\.config\konv
// Linux: ~/.config/konv
// MacOS: ~/.config/konv
func selfBase() string {
	dir, _ := os.UserHomeDir()
	return path.Join(dir, ".config", "konv")
}

func getPath(target, ext string, isSelf bool) string {
	var b string
	if isSelf {
		b = selfBase()
	} else {
		b = base()
	}
	return path.Join(b, target+"."+ext)
}

// GetConvPath Get Conv Rule Path
func GetConvPath(target string, isSelf bool) string {
	return getPath(target, "csv", isSelf)
}

// GetConvCfgPath Get Conv Cfg Path
func GetConvCfgPath(target string, isSelf bool) string {
	return getPath(target, "cfg", isSelf)
}
