package utils

import (
	"code.google.com/p/gcfg"
	"os/user"
)

type _config struct {
	Database struct {
		Name string
	}
}

var Config _config

func init() {
	usr, _ := user.Current()

	file := usr.HomeDir + "/.todoin"

	gcfg.ReadFileInto(&Config, file)
}
