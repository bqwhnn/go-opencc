package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

const (
	config = "s2t"
	suffix = ""
)

var (
	dir   string
	count int
	paths = []string{
		"F:/example",
	}
)

func main() {
	var err error
	dir, err = os.Getwd()
	if err != nil {
		fmt.Println(dir, err)
	}

	if len(os.Args) < 2 {
		return
	}

	filepaths := []string{}
	if os.Args[1] == "server" {
		filepaths = paths
	} else {
		filepaths = os.Args[1:]
	}

	for _, path := range filepaths {
		err = filepath.Walk(path, convert)
		if err != nil {
			fmt.Printf("filepath.Walk error %v\n", err)
		}
	}

	fmt.Printf("total file count : %d\n", count)

	time.Sleep(5 * time.Second)
}

func convert(fp string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}

	if f.IsDir() {
		return nil
	}

	if suffix != "" && filepath.Ext(fp) != suffix {
		return nil
	}

	fmt.Println(fp)

	command := dir + "/opencc-1.0.4/bin/opencc.exe -i " + fp + " -o " + fp + " -c " + dir + "/opencc-1.0.4/share/opencc/" + config + ".json"
	cmd := exec.Command("cmd.exe", "/c", "start "+command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		return err
	}

	count = count + 1

	return nil
}
