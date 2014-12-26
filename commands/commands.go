package commands

import (
	"flag"
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
)

var appName string
var containers = make(map[string]*Command)

type Command struct {
	Usage func() string
	Run   func(args []string) int
}

type Args struct {
	Flag      *flag.FlagSet
	Verbose   *bool
	Separator *string
	Task      map[string]*string
}

func Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <commands> <options>\n", appName)
	os.Exit(255)
}

func Register(name string, cmd *Command) {
	containers[name] = cmd
}

func Run() {
	if len(os.Args) < 2 {
		Usage()
		os.Exit(0)
	}

	// check table
	if !db.TableExist("task") && os.Args[1] != "init" {
		fmt.Println("Table not found, please init")
		os.Exit(255)
	}

	exit := 255
	name := os.Args[1]
	if cmd, ok := Exist(name); ok {
		exit = cmd.Run(os.Args[2:])
	}
	os.Exit(exit)
}

func Exist(name string) (*Command, bool) {
	cmd, ok := containers[name]
	return cmd, ok
}

func Init() {
	appName = os.Args[0]
}

func parseFlag(name string) *Args {
	args := new(Args)
	args.Flag = flag.NewFlagSet("get", flag.ContinueOnError)
	args.Flag.Init(name, flag.ContinueOnError)
	args.Verbose = args.Flag.Bool("verbose", false, "")
	args.Separator = args.Flag.String("separator", "  ", "")

	args.Task = make(map[string]*string)
	args.Task["assignby"] = args.Flag.String("assignby", "", "")
	args.Task["assignto"] = args.Flag.String("assignto", "", "")
	args.Task["status"] = args.Flag.String("status", "", "")
	args.Task["value"] = args.Flag.String("value", "", "")
	args.Task["duedate"] = args.Flag.String("duedate", "", "")

	return args
}
