package param

import (
	"github.com/KevinZonda/apt-pac/pkg/utils"
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

func Parse(s string) Mod {
	m := dft()
	if !strings.HasPrefix(s, ":") {
		return m
	}
	s = s[1:]
	if s == "y" {
		m.SkipConfirm = true
	}
	return m
}

func CleanArg(s string) string {
	ss := strings.Split(s, "/")
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
