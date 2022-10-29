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
	_, to, dec, err := loader.Load("/home/kevin/Desktop/apt-pac/conv.csv")

	// fmt.Printf("from: %s\nto: %s\ndec: %+v\nerr: %v\n", from, to, dec, err)
	if err != nil {
		fmt.Println("load conv law failed!")
		return
	}

	pattern, vars, ok := loader.Conv(dec, os.Args[1:])
	if !ok {
		fmt.Println("conv law failed!")
		return
	}
	argses := loader.PatternToArgs(pattern, vars)
	scanner := bufio.NewReader(os.Stdin)
	for _, arg := range argses {
		fmt.Printf("%+v\n", arg)
		fmt.Print("Continue [y/n]?")
		ans, _ := scanner.ReadString('\n')
		if utils.Trim(ans) != "y" {
			fmt.Println("Abort")
			return
		}
		continue

		cmd := exec.Command(to)

		cmd.Args = arg

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		_ = cmd.Start()
	}
}
