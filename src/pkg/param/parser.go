package param

import (
	"github.com/KevinZonda/konv/pkg/utils"
	"github.com/KevinZonda/konv/pkg/verbose"
	"strings"
)

type Mod struct {
	Ok          bool
	SkipConfirm bool
	ShowCmdList bool
	Verbose     bool
}

func dft() Mod {
	return Mod{
		Ok:          false,
		SkipConfirm: false,
		ShowCmdList: false,
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
		case "y", "sc":
			m.SkipConfirm = true
			break
		case "c":
			m.SkipConfirm = false
			break
		case "l":
			m.ShowCmdList = true
			break
		case "v":
			m.Verbose = true
			verbose.VerboseMode = true
			verbose.PrintQueue()
		}
	}
	m.Ok = true
	return m
}

func GetSourceFromArg0(s string) string {
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
