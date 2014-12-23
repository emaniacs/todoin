package main

import (
	"github.com/emaniacs/todoin/commands"
)

func init() {
	commands.Init()
}

func main() {
	commands.Run()
}
