package utils

import (
	"os"
	"os/exec"
)

func Gomod(arg ...string) {
	if len(arg) == 0 {
		panic("")
	} else {
		cmd := exec.Command(os.Getenv("GOPATH")+"\\bin\\go.exe", getArgs("mod", arg...)...)
		err := cmd.Run()
		Panic(err)
	}
}

func getArgs(command string, args ...string) (fargs []string) {
	fargs = append(fargs, command)
	fargs = append(fargs, args...)
	return fargs
}

func MkDirs(dirPath string) {
	err := os.MkdirAll(dirPath, 0666)
	Panic(err)
}
