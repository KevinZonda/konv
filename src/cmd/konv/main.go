package main

import (
	"fmt"
	"github.com/KevinZonda/konv/pkg/console"
	"github.com/KevinZonda/konv/pkg/loader"
	"github.com/KevinZonda/konv/pkg/param"
	"github.com/KevinZonda/konv/pkg/path"
	"github.com/KevinZonda/konv/pkg/process"
	"github.com/KevinZonda/konv/pkg/utils"
	"os"
)

func main() {
	arg1 := utils.Trim(param.CleanArg((os.Args[0])))
	if arg1 == "" {
		panic("Cannot found target rule")
	}

	_, to, dec, err := loader.Load(path.GetConversionPath(arg1))
	utils.PanicIfNotNil(err, "load conv law failed!")

	osArgs := os.Args[1:]

	if len(osArgs) == 0 {
		fmt.Println(arg1)
		fmt.Println("powered by konv")
		return
	}

	cfg := param.Parse(osArgs[0])
	if cfg.Ok {
		osArgs = osArgs[1:]
	}

	pattern, vars, ok := loader.Conv(dec, utils.TrimAll(osArgs))
	if !ok {
		panic("Parse failed. Please confirm that your command is correct!")
	}
	argses := loader.PatternToArgs(pattern, vars)
	runs := parseArgs(to, argses)

	if !cfg.SkipConfirm {
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

	process.Runs(runs)
}
