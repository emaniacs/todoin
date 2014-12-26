package commands

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/emaniacs/todoin/db"
	"github.com/emaniacs/todoin/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var regexString = ".*TODO: +([^@$?!]*)([@$?].*[^ ])?"
var Ask, Insert *bool
var fchan = make(chan string)
var tchan = make(chan *db.Task)
var Exts []string
var Wg sync.WaitGroup

func parseFile() {
	re := regexp.MustCompile(regexString)
	for {
		fn := <-fchan
		file, err := os.Open(fn)
		if err != nil {
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		line := 1
		for scanner.Scan() {
			text := scanner.Text()
			found := re.FindStringSubmatch(text)
			if len(found) > 0 {
				Wg.Add(1)
				tchan <- generateTask(found, fn, line)
			}
			line++
		}
	}
}

func addTask() {
	for {
		t := <-tchan
		if *Insert {
			fmt.Fprintf(os.Stdout, "[%s:%d \"%s\"] ", t.Filename, t.Line, t.Value)
			insert := true
			if *Ask {
				var yes string
				fmt.Print("Insert (y/n): ")
				fmt.Scanf("%s", &yes)
				if strings.Trim(yes, " ") != "y" {
					insert = false
				}
			} else {
				fmt.Println("")
			}

			if insert {
				key := db.Insert(t)
				fmt.Printf(" -> %d\n", key)
			} else {
				fmt.Println(" -> No")
			}
		} else {
			fmt.Printf("%s:%s \"%s\"\n", t.Filename, t.Line, t.Value)
		}
		Wg.Done()
	}
}

func generateTask(found []string, fn string, line int) *db.Task {
	value := found[1]

	task := new(db.Task)
	task.Value = value
	task.Filename = fn
	task.Line = line
	if len(found) > 2 {
		for _, val := range strings.Split(found[2], " ") {
			if status, ok := utils.IsDone(val); ok {
				task.Status = status
			} else if utils.IsAssignBy(val) {
				task.AssignBy = val[1:]
			} else if utils.IsAssignTo(val) {
				task.AssignTo = val[1:]
			} else if utils.IsDueDate(val) {
				task.DueDate = val[1:]
			}
		}
	}
	return task
}

func prefixMatch(file string) bool {
	for _, val := range Exts {
		if strings.HasSuffix(file, val) {
			return true
		}
	}
	return false
}

func init() {
	Register("gen", &Command{
		Usage: func() string {
			return fmt.Sprintf(`Generate task based on TODO text in a file
Usage:
	%s gen <options>
Options:
	-path=dir	Path of file (required)
	-ext=ext	Extension file (required)
	-insert		Force insert (default is true)
	-ask		Ask before insert (default is false)
Example:
	$ %s -path=/tmp -ext=*.go 
	$ %s -path=/tmp,/home -ext=*.go,*.php -ask
			`, appName, appName, appName)
		},
		Run: func(args []string) int {
			flg := flag.NewFlagSet("get", flag.ContinueOnError)
			flg.Init("gen", flag.ContinueOnError)
			paths := flg.String("path", "", "")
			exts := flg.String("ext", "", "")
			Insert = flg.Bool("insert", true, "")
			Ask = flg.Bool("ask", false, "")
			flg.Parse(args)

			if len(*paths) < 1 || len(*exts) < 1 {
				fmt.Fprintln(os.Stderr, "Not enough argument")
				return 255
			}

			for _, pre := range strings.Split(*exts, ",") {
				Exts = append(Exts, strings.Replace(pre, "*", "", -1))
			}

			go addTask()
			go parseFile()
			for _, path := range strings.Split(*paths, ",") {
				Wg.Add(1)
				go func() {
					filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
						if !f.IsDir() && prefixMatch(path) {
							fchan <- path
						}
						return nil
					})
					Wg.Done()
				}()
			}
			Wg.Wait()

			return 255
		},
	})
}
