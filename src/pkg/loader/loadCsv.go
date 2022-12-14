package loader

import (
	"errors"
	"strings"

	"github.com/KevinZonda/konv/pkg/csv"
	"github.com/KevinZonda/konv/pkg/decision"
	"github.com/KevinZonda/konv/pkg/io"
	"github.com/KevinZonda/konv/pkg/utils"
)

func loadFromToCsv(path string) (from, to string, result map[string]string, err error) {
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

	result = csv.ParseLines(lines[1:], ",")
    
	return
}

func parseCsvContentToTree(head string, result map[string]string) *decision.CheckTree {
	f := &decision.CheckTree{
		Key: head,
	}
	for k, v := range result {
		appendCommandToTree(k, f, v)
	}
	return f
}

func LoadConvRule(path string) (from, to string, dec *decision.CheckTree, err error) {
	var content map[string]string
	from, to, content, err = loadFromToCsv(path)
	if err != nil {
		return
	}
	dec = parseCsvContentToTree(from, content)
	return
}
