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
	ss := utils.TrimAll(strings.Split(s, ":"))
	for _, p := range ss {
		switch p {
		case "y":
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
