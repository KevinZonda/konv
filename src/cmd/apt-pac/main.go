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

	if err != nil {
		fmt.Println("load conv law failed!")
		return
	}

	pattern, vars, ok := loader.Conv(dec, utils.TrimAll(os.Args[1:]))
	if !ok {
		fmt.Println("Parse failed. Please confirm that your command is correct!")
		return
	}
	argses := loader.PatternToArgs(pattern, vars)

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

	for _, arg := range argses {
		cmd := exec.Command(to, arg...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		_ = cmd.Start()
		_ = cmd.Wait()
	}
}
