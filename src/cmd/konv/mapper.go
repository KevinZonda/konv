package main

import (
	"github.com/KevinZonda/konv/pkg/process"
	"strings"
)

func mapToRunnableList(name string, argsList [][]string, env []string) []process.Runable {
	var list []process.Runable
	if len(env) == 0 {
		env = nil
	}
	for _, v := range argsList {
		if strings.HasPrefix(v[0], "@") {
			list = append(list, process.Runable{
				Name: v[0][1:],
				Args: v[1:],
				Env:  env,
			})
			continue
		}
		list = append(list, process.Runable{
			Name: name,
			Args: v,
		})
	}
	return list
}
