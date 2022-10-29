package io

import (
	"os"
	"strings"
)

func ReadAllText(path string) (dat string, err error) {
	var datB []byte
	datB, err = os.ReadFile(path)
	if err != nil {
		return
	}
	dat = string(datB)
	return
}

func ReadAllLines(path string) (lines []string, err error) {
	var text string
	text, err = ReadAllText(path)
	if err != nil {
		return
	}
	lines = strings.Split(text, "\n")
	return
}
