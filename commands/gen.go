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
)

var regexString = ".*TODO: +([^@$?!]*)([@$?].*[^ ])?"

func parseFile(fchan chan string) {
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
			task := re.FindStringSubmatch(text)
			if len(task) > 0 {
				value := task[1]
				other := ""
				if len(task) > 2 {
					other = task[2]
				}
				key := addTask(value, other)
				fmt.Printf("(%d) %s:%s \"%s\"\n", key, fn, line, value)
			}
			line++
		}
	}
}

func parseDir(dirs string, prefixs string, fchan chan string) {
	for _, dir := range strings.Split(dirs, ",") {
		dir = strings.TrimRight(dir, "/")
		dirStat, err := os.Stat(dir)
		if err != nil {
			continue
		}

		if dirStat.IsDir() {
			for _, prefix := range strings.Split(prefixs, ",") {
				files, _ := filepath.Glob(dir + "/" + prefix)
				for _, file := range files {
					fchan <- file
				}
			}
		} else {
			fchan <- dir
		}
	}
}

func addTask(value string, other string) int64 {
	task := new(db.Task)
	task.Value = value
	for _, val := range strings.Split(other, " ") {
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

	return db.Insert(task)
}

func init() {
	Register("gen", &Command{
		Usage: func() string {
			return "Usage of gen"
		},
		Run: func(args []string) int {
			flg := flag.NewFlagSet("get", flag.ContinueOnError)
			flg.Init("gen", flag.ContinueOnError)
			paths := flg.String("path", "", "")
			files := flg.String("file", "", "")
			flg.Parse(args)

			if len(*paths) < 1 || len(*files) < 1 {
				fmt.Fprintln(os.Stderr, "Not enough argument")
				return 255
			}

			fchan := make(chan string)
			go parseFile(fchan)
			parseDir(*paths, *files, fchan)

			return 255
		},
	})
}
