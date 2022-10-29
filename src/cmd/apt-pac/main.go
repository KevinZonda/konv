package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hi")

	cmd := exec.Command("pacman")

	cmd.Args = os.Args

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	_ = cmd.Start()
}
