package main

import (
	"fmt"
	"github.com/emaniacs/todoin/commands"
	"os"
)

func main() {
	var exit int
	var msg string
	if len(os.Args) < 2 {
		exit, msg = commands.Help()
		fmt.Println(msg)
		os.Exit(exit)
	}

	switch os.Args[1] {
	case "read", "r":
		exit, msg = commands.Read()
	case "add", "a":
		exit, msg = commands.Add()
	default:
		exit = -1
		msg = "Unknown command \"" + os.Args[1] + "\""
	}

	if msg != "" {
		fmt.Println(msg)
	}
	os.Exit(exit)
}
