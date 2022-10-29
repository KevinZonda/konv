package main

import (
	"fmt"
	"github.com/KevinZonda/apt-pac/pkg/loader"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hi")

	from, to, dec, err := loader.Load("/home/kevin/Desktop/apt-pac/conv.csv")

	fmt.Printf("from: %s\nto: %s\ndec: %+v\nerr: %v", from, to, dec, err)
	return

	cmd := exec.Command("pacman")

	cmd.Args = os.Args

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	_ = cmd.Start()
}
