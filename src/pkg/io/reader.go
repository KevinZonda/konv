package io

import "os"

func ReadAllText(path string) (dat string, err error) {
	var datB []byte
	datB, err = os.ReadFile(path)
	if err != nil {
		return
	}
	dat = string(datB)
	return
}
