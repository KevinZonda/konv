package csv

import (
	"strings"

	"github.com/KevinZonda/konv/pkg/utils"
)

func ParseLines(lines []string, sep string) map[string]string {
	pairs := make(map[string]string)
	if lines == nil || len(lines) < 1 {
		return pairs
	}
	for _, line := range lines {
		trimed := utils.Trim(line)
		if line == "" {
			continue
		}
		kvp := strings.Split(trimed, sep)
		if len(kvp) != 2 {
			continue
		}
		pairs[utils.Trim(kvp[0])] = utils.Trim(kvp[1])
	}
	return pairs
}

func Parse(c string, sep string) map[string]string {
	return ParseLines(strings.Split(c, "\n"), sep)
}
