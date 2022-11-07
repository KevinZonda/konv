package cfg

import (
	"github.com/KevinZonda/konv/pkg/utils"
	"strings"
)

func ParseLines(lines []string) map[string]string {
	pairs := make(map[string]string)
	if lines == nil || len(lines) < 1 {
		return pairs
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
		pairs[utils.Trim(kvp[0])] = utils.Trim(kvp[1])
	}
	return pairs
}
