package main

import (
	"fmt"
	"github.com/KevinZonda/konv/pkg/console"
	"github.com/KevinZonda/konv/pkg/param"
	"github.com/KevinZonda/konv/pkg/process"
)

func run(runCfg param.Mod, runs []process.Runable) {
	if !runCfg.SkipConfirm {
		fmt.Println("Please ensure following commands are correct!")
		for i, r := range runs {
			fmt.Printf("%d: %s %+v\n", i, r.Name, r.Args)
		}
		fmt.Print("Continue [y/n]?")
		isOk := console.ReadYes()
		if !isOk {
			fmt.Println("Abort")
			return
		}
	}

	if runCfg.ShowCmdList && runCfg.SkipConfirm {
		for i, r := range runs {
			fmt.Printf("%d: %s %+v\n", i, r.Name, r.Args)
		}
	}

	process.Runs(runs, func(r process.Runable, err error) bool {
		fmt.Printf(
			"run failed:\n"+
				"cmd: %s %+v\n"+
				"err: %+v\n", r.Name, r.Args, err)
		fmt.Print("Continue [y/n?")
		return console.ReadYes()
	})
}
