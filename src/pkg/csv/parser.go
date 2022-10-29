package csv

import "strings"

func Parse(c string, sep string) map[string]string {
	lines := strings.Split(c, "\n")
	pairs := make(map[string]string)
	for _, line := range lines {
		trimed := trim(line)
		kvp := strings.Split(trimed, sep)
		if len(kvp) != 2 {
			continue
		}
		pairs[trim(kvp[0])] = trim(kvp[1])
	}
	return pairs
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
