package csv

import "strings"

func ParseLines(lines []string, sep string) map[string]string {
	pairs := make(map[string]string)
	if lines == nil || len(lines) < 1 {
		return pairs
	}
	for _, line := range lines {
		trimed := trim(line)
		if line == "" {
			continue
		}
		kvp := strings.Split(trimed, sep)
		if len(kvp) != 2 {
			continue
		}
		pairs[trim(kvp[0])] = trim(kvp[1])
	}
	return pairs
}

func Parse(c string, sep string) map[string]string {
	return ParseLines(strings.Split(c, "\n"), sep)
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
