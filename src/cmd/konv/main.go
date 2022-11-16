package main

import (
	"fmt"
	"github.com/KevinZonda/konv/pkg/cfg"
	"github.com/KevinZonda/konv/pkg/loader"
	"github.com/KevinZonda/konv/pkg/param"
	"github.com/KevinZonda/konv/pkg/path"
	"github.com/KevinZonda/konv/pkg/utils"
	"os"
)

func main() {
	source := utils.Trim(param.GetSourceFromArg0(os.Args[0]))
	if source == "" {
		panic("Cannot found target rule")
	}

	isSelf := true
	_, to, dec, err := loader.Load(path.GetConvPath(source, true))
	if err != nil {
		isSelf = false
		_, to, dec, err = loader.Load(path.GetConvPath(source, isSelf))
	}
	utils.PanicIfNotNil(err, "load conv law failed!")

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print(source)
		fmt.Println(" - powered by konv")
		return
	}

	runCfg := param.Parse(args[0])
	ruleCfg := cfg.GetConfig(path.GetConvCfgPath(source, isSelf))

	if runCfg.Ok {
		args = args[1:]
	} else {
		runCfg = param.Parse(ruleCfg.Param)
	}

	pattern, vars, ok := loader.Conv(dec, utils.TrimAll(args))
	if !ok {
		panic("Parse failed. Please confirm that your command is correct!")
	}
	argses := loader.PatternToArgs(pattern, vars)
	runs := mapToRunnableList(to, argses, ruleCfg.Env)
	run(runCfg, runs)
}
