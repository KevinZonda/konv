package main

import (
	"bufio"
	"fmt"
	"github.com/KevinZonda/apt-pac/pkg/loader"
	"github.com/KevinZonda/apt-pac/pkg/utils"
	"os"
	"os/exec"
	"strings"
)

func cleanArg(s string) string {
	ss := strings.Split(s, "/")
	s = ss[len(ss)-1]
	sb := strings.Builder{}
	for _, c := range s {
		if !isLetter(c) {
			sb.Reset()
			continue
		}
		sb.WriteRune(c)
	}

	return sb.String()
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func main() {
	arg1 := utils.Trim(cleanArg(os.Args[0]))
	if arg1 == "" {
		panic("Cannot found target rule")
	}

	_, to, dec, err := loader.Load("/etc/konv/" + arg1 + ".csv")
	utils.PanicIfNotNil(err, "load conv law failed!")

	osArgs := os.Args[1:]

	needConfirm := true

	switch len(osArgs) {
	case 0:
		fmt.Println("konv")
		return
	default:
		if osArgs[0] == "y" {
			needConfirm = false
			osArgs = osArgs[1:]
		}
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
