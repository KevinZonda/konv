package main

import "fmt"

func logo(from string) {
	fmt.Println()
	fmt.Println("Welcome to Konv by KevinZonda")
	fmt.Println("-----------------------------")
	fmt.Println("Version: " + VERSION)
	fmt.Println("Trigger: " + from)

	fmt.Println()

	fmt.Println("-- Empty action list. No action has been done!")
}

const VERSION = "1.0 - Super Kevin"
