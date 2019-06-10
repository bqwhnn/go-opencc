package main

import (
    "os"
    "fmt"
    "os/exec"
    "path/filepath"
	"syscall"
	"time"
)

const (
	config = "s2t"
	in = "test"
	suffix = ""
)

var (
	dir string
)

func main() {
	var err error
	dir, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path := dir + "/" + in

	err = filepath.Walk(path, convert)
	if err != nil {
		fmt.Printf("filepath.Walk error %v\n", err)
	}

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
	cmd := exec.Command("cmd.exe", "/c", "start " + command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow:true}
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}