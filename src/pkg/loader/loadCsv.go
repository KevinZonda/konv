package loader

import (
	"errors"
	"strings"

	"github.com/KevinZonda/apt-pac/pkg/csv"
	"github.com/KevinZonda/apt-pac/pkg/io"
	"github.com/KevinZonda/apt-pac/pkg/utils"
)

func LoadFromToStyleCsv(path string) (from, to string, result map[string]string, err error) {
	var lines []string
	lines, err = io.ReadAllLines(path)
	if err != nil {
		return
	}
	if len(lines) <= 1 {
		err = errors.New("not enough lines")
		return
	}

	firstLine := utils.Trim(lines[0])
	headers := strings.Split(firstLine, ",")
	if len(headers) != 2 {
		err = errors.New("not enough lines")
		return
	}
	from, to = headers[0], headers[1]

	result = csv.ParseLines(lines[1:], "")
	return
}
