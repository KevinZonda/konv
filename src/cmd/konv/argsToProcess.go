package main

import (
	"github.com/KevinZonda/apt-pac/pkg/process"
	"strings"
)

func parseArgs(name string, argses [][]string) []process.Runable {
	var list []process.Runable
	for _, v := range argses {
		if strings.HasPrefix(v[0], "@") {
			list = append(list, process.Runable{
				Name: v[0][1:],
				Args: v[1:],
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
