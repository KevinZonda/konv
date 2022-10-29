package utils

import "strings"

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func TrimAll(ss []string) []string {
	var rst []string
	for _, s := range ss {
		t := Trim(s)
		if t != "" {
			rst = append(rst, t)
		}
	}
	return rst
}
