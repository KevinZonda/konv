package main

import (
	"bufio"
	"fmt"
	"github.com/KevinZonda/apt-pac/pkg/loader"
	"github.com/KevinZonda/apt-pac/pkg/utils"
	"os"
	"os/exec"
)

func main() {
	_, to, dec, err := loader.Load("/etc/apt-pac/rule.csv")
	utils.PanicIfNotNil(err, "load conv law failed!")

	osArgs := os.Args[1:]

	needConfirm := true

	switch len(osArgs) {
	case 0:
		fmt.Println("apt-pac")
		return
	default:
		if osArgs[0] == "y" {
			needConfirm = false
		}
		osArgs = osArgs[1:]
	}

	pattern, vars, ok := loader.Conv(dec, utils.TrimAll(osArgs))
	if !ok {
		panic("Parse failed. Please confirm that your command is correct!")
	}
	argses := loader.PatternToArgs(pattern, vars)

	if needConfirm {
		scanner := bufio.NewReader(os.Stdin)
		fmt.Println("Please ensure following commands are correct!")
		for i, arg := range argses {
			fmt.Printf("%d: %s %+v\n", i, to, arg)
		}
		fmt.Print("Continue [y/n]?")
		ans, _ := scanner.ReadString('\n')
		if utils.Trim(ans) != "y" {
			fmt.Println("Abort")
			return
		}
	}

	for _, arg := range argses {
		cmd := exec.Command(to, arg...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		_ = cmd.Start()
		_ = cmd.Wait()
	}
}
