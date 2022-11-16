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
	for _, line := range lines {
		trimed := utils.Trim(line)
		if line == "" {
			continue
		}
		kvp := strings.SplitN(trimed, "=", 2)
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
