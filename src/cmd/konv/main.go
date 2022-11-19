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

	ruleCfg := cfg.GetConfig(path.GetConvCfgPath(source, isSelf))
	runCfg := param.Mod{
		Ok: false,
	}
	if len(args) > 0 {
		runCfg = param.Parse(args[0])
		if runCfg.Ok {
			args = args[1:]
		}
	}
	if !runCfg.Ok {
		runCfg = param.Parse(ruleCfg.Param)
	}

	pattern, vars, ok := loader.GetDecisionResult(dec, utils.TrimAll(args))
	if !ok {
		panic("Parse failed. Please confirm that your command is correct!")
	}
	cmdsArgs := loader.CombineToCommandArgs(pattern, vars)
	if len(cmdsArgs) == 0 {
		fmt.Printf("%s - powered by konv\n", source)
		fmt.Println("-- Empty action list. No oper has been done!")
		return
	}
	runs := mapToRunnableList(to, cmdsArgs, ruleCfg.Env)
	run(runCfg, runs)
}
