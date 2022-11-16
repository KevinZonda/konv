package cfg

import (
	"github.com/KevinZonda/konv/pkg/utils"
	"strings"
)

type Mod struct {
	Param string
}

func ParseLines(lines []string) Mod {
	m := Mod{
		Param: "",
	}
	if lines == nil || len(lines) < 1 {
		return m
	}
	lines = utils.TrimAll(lines)
	for _, line := range lines {
		kvp := strings.SplitN(line, "=", 2)
		if len(kvp) != 2 {
			continue
		}
		if utils.Trim(kvp[0]) == "param" {
			m.Param = utils.Trim(kvp[1])
			continue
		}
	}
	return m
}
