package loader

import (
	"github.com/KevinZonda/apt-pac/pkg/decision"
	"github.com/KevinZonda/apt-pac/pkg/utils"
	"strings"
)

func Conv(r *decision.CheckTree, args []string) (pattern string, vars []string, ok bool) {
	curr := r
	isV := false
	for _, arg := range args {
		if curr == nil {
			ok = false
			return
		}
		curr, isV = curr.Next(arg)
		if isV {
			vars = append(vars, arg)
		}
	}
	ok = true
	if curr == nil || curr.Data == nil {
		ok = false
		return
	}
	pattern = *curr.Data
	ok = true
	return
}

func PatternToArgs(pattern string, vars []string) [][]string {
	cmds := utils.TrimAll(strings.Split(pattern, ";"))
	var rst [][]string
	for _, cmd := range cmds {
		finalArgs := []string{}
		args := utils.TrimAll(strings.Split(cmd, " "))
		for _, arg := range args {
			switch arg {
			case "$":
				if len(vars) != 1 {
					panic("wrong len")
				}
			case "$$":
				finalArgs = append(finalArgs, vars...)
			default:
				finalArgs = append(finalArgs, arg)
			}
		}
		rst = append(rst, finalArgs)
	}
	return rst
}
