package main

import (
	"fmt"
	"github.com/emaniacs/todoin/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		exit, help := commands.Help()
		fmt.Println(help)
		os.Exit(exit)
	}

	switch os.Args[1] {
	case "show", "s":
		var exit int
		msg := []string{}
		exit, msg = commands.Show()
		for key := range msg {
			fmt.Println(msg[key])
		}
		os.Exit(exit)
	case "add", "a":
		exit, msg := commands.Add()
		fmt.Println(msg)
		os.Exit(exit)
	default:
		fmt.Println("Unknown command \"" + os.Args[1] + "\"")
		os.Exit(-1)
	}
}
