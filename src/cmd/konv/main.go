package main

import (
	"fmt"
	"github.com/KevinZonda/apt-pac/pkg/console"
	"github.com/KevinZonda/apt-pac/pkg/loader"
	"github.com/KevinZonda/apt-pac/pkg/param"
	"github.com/KevinZonda/apt-pac/pkg/process"
	"github.com/KevinZonda/apt-pac/pkg/utils"
	"os"
)

func main() {
	arg1 := utils.Trim(param.CleanArg((os.Args[0])))
	if arg1 == "" {
		panic("Cannot found target rule")
	}

	_, to, dec, err := loader.Load("/etc/konv/" + arg1 + ".csv")
	utils.PanicIfNotNil(err, "load conv law failed!")

	osArgs := os.Args[1:]

	var cfg param.Mod

	switch len(osArgs) {
	case 0:
		fmt.Println("konv")
		return
	default:
		cfg = param.Parse(osArgs[0])
		if cfg.Ok {
			osArgs = osArgs[1:]
		}
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
