package utils

import (
	"os"
	"os/exec"
)

func Gomod(arg ...string) {
	if len(arg) == 0 {
		panic("")
	} else if len(arg) == 1 {
		cmd := exec.Command(os.Getenv("GOPATH")+"\\bin\\go.exe", "mod", arg[0])
		err := cmd.Run()
		Panic(err)
	} else if len(arg) == 2 {
		cmd := exec.Command(os.Getenv("GOPATH")+"\\bin\\go.exe", "mod", arg[0], arg[1])
		err := cmd.Run()
		Panic(err)
	}
}

func MkDirs(dirPath string) {
	err := os.MkdirAll(dirPath, 0666)
	Panic(err)
}
