package csv

import (
	"strings"

	"github.com/KevinZonda/konv/pkg/utils"
)

const COMMENT_PREFIX = "#"

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
        if strings.HasPrefix(trimed, COMMENT_PREFIX) {
			continue
		}
		kvp := strings.SplitN(trimed, sep, 2)
		if len(kvp) != 2 {
			continue
		}
		pairs[utils.Trim(kvp[0])] = utils.Trim(kvp[1])
	}
	return pairs
}
