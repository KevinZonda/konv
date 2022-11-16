package cfg

import (
	"github.com/KevinZonda/konv/pkg/io"
	"github.com/KevinZonda/konv/pkg/utils"
	"strings"
)

type Mod struct {
	Param string
	Env   []string
}

func ParseLines(lines []string) Mod {
	m := dft()
	if lines == nil || len(lines) < 1 {
		return m
	}
	lines = utils.TrimAll(lines)
	for _, line := range lines {
		kvp := strings.SplitN(line, "=", 2)
		if len(kvp) != 2 {
			continue
		}
		key := utils.Trim(kvp[0])
		value := utils.Trim(kvp[1])
		switch key {
		case "param":
			m.Param = value
			break
		case "env":
			m.Env = append(m.Env, key)
			break
		}
	}
	return m
}

func dft() Mod {
	return Mod{
		Param: "",
		Env:   nil,
	}
}

func GetConfig(path string) (result Mod) {
	s, err := io.ReadAllLines(path)
	if err != nil {
		return dft()
	}
	return ParseLines(s)
}
