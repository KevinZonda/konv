package param

import (
	"github.com/KevinZonda/konv/pkg/cfg"
	"github.com/KevinZonda/konv/pkg/io"
	"github.com/KevinZonda/konv/pkg/utils"
	"strings"
)

type Mod struct {
	Ok          bool
	SkipConfirm bool
}

func dft() Mod {
	return Mod{
		Ok:          false,
		SkipConfirm: false,
	}
}

func ParseFromFile(path string) (isOk bool, mod Mod) {
	mod = dft()
	s, err := io.ReadAllLines(path)
	if err != nil {
		return false, mod
	}
	m := cfg.ParseLines(s)
	param, ok := m["param"]
	if !ok {
		return false, mod
	}
	return true, Parse(param)
}

func Parse(s string) Mod {
	m := dft()
	if !strings.HasPrefix(s, ":") {
		return m
	}
	ss := utils.TrimAll(strings.Split(s, ":"))
	for _, p := range ss {
		switch p {
		case "y", "sc":
			m.SkipConfirm = true
			break
		case "c":
			m.SkipConfirm = false
			break
		}
	}
	return m
}

func CleanArg(s string) string {
	ss := strings.Split(s, "/")
	s = ss[len(ss)-1]
	ss = strings.Split(s, "\\")
	s = ss[len(ss)-1]
	sb := strings.Builder{}
	for _, c := range s {
		if !utils.IsLetter(c) {
			sb.Reset()
			continue
		}
		sb.WriteRune(c)
	}

	return sb.String()
}
