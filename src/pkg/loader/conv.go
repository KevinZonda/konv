package loader

import (
	"github.com/KevinZonda/konv/pkg/decision"
	"github.com/KevinZonda/konv/pkg/utils"
	"strings"
)

func GetDecisionResult(r *decision.CheckTree, args []string) (pattern string, vars []string, ok bool) {
	curr := r
	if len(args) == 0 {
		if curr.Data != nil {
			return *curr.Data, nil, true
		}
		isWildcard := false
		curr, isWildcard = curr.Next("")
		if isWildcard {
			return "", nil, false
		}
		if curr.Data != nil {
			return *curr.Data, nil, true
		}
	}

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

func CombineToCommandArgs(pattern string, vars []string) [][]string {
	cmds := utils.TrimAll(strings.Split(pattern, ";"))
	var rst [][]string
	for _, cmd := range cmds {
		finalArgs := []string{}
		args := utils.TrimAll(strings.Split(cmd, " "))
		for _, arg := range args {
			if arg == "$" {
				if len(vars) == 0 {
					panic("wrong len")
				}
				finalArgs = append(finalArgs, vars[0])
				vars = vars[1:]
				continue
			}
			if arg == "$$" {
				finalArgs = append(finalArgs, vars...)
				continue
			}
			finalArgs = append(finalArgs, arg)
		}
		rst = append(rst, finalArgs)
	}
	return rst
}
