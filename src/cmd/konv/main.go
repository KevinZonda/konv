package main

import (
	"os"

	"github.com/KevinZonda/konv/pkg/cfg"
	"github.com/KevinZonda/konv/pkg/decision"
	"github.com/KevinZonda/konv/pkg/loader"
	"github.com/KevinZonda/konv/pkg/param"
	"github.com/KevinZonda/konv/pkg/path"
	"github.com/KevinZonda/konv/pkg/utils"
	"github.com/KevinZonda/konv/pkg/verbose"
)

func loadConvRule(source string) (to string, dec *decision.CheckTree, isSelf bool) {
	convPath := path.GetConvPath(source, true)
	verbose.PrintF("load conv %s", convPath)
	_, to, dec, err := loader.LoadConvRule(convPath)
	if err == nil {
		return to, dec, true
	}
	convPath = path.GetConvPath(source, false)
	verbose.PrintF("load failed, try %s", convPath)
	_, to, dec, err = loader.LoadConvRule(convPath)
	if err != nil {
		panic("load conv law failed!")
	}
	return to, dec, false
}

func main() {
	source := utils.Trim(param.GetSourceFromArg0(os.Args[0]))

	if source == "" {
		panic("Cannot found target rule")
	}
	if source == "konv" {
		logo(source)
		return
	}

	to, dec, isSelf := loadConvRule(source)

	args := os.Args[1:]

	ruleCfg := cfg.GetConfig(path.GetConvCfgPath(source, isSelf))
	runCfg := param.Mod{
		Ok: false,
	}

	if len(args) > 0 {
		verbose.PrintF("try parse arg[1] to get cfg")
		runCfg = param.Parse(args[0])
		if runCfg.Ok {
			args = args[1:]
		}
	}
	if !runCfg.Ok {
		verbose.PrintF("try parse rule config to get run cfg")
		runCfg = param.Parse(ruleCfg.Param)
	}

	pattern, vars, ok := loader.GetDecisionResult(dec, args)

	if !ok {
		if len(args) == 0 {
			logo(source)
			return
		}
		panic("Parse failed. Please confirm that your command is correct!")
	}
	cmdsArgs := loader.CombineToCommandArgs(pattern, vars)
	if len(cmdsArgs) == 0 {
		logo(source)
		return
	}
	runs := mapToRunnableList(to, cmdsArgs, ruleCfg.Env)
	run(runCfg, runs)
}
