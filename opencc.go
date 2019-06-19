package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const (
	config = "s2t"
	suffix = ""
)

var (
	dir         string
	serverPaths = []string{
		"E:/git/go-opencc/test",
	}
	filetype    = []string{".lua", ".json", ".ts", ".txt"}
	replaceText = map[string]string{
		"首冲":  "首儲",
		"邮件":  "信箱",
		"充值":  "儲值",
		"点击":  "點選",
		"激活码": "禮品碼",
		"变强":  "成長",
		"帮派":  "公會",
		"登录":  "登入",
		"登陆":  "登入",
		"确定":  "確認",
		"服务器": "伺服器",
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

	paths := []string{}
	if os.Args[1] == "server" {
		paths = serverPaths
	} else {
		paths = os.Args[1:]
	}

	start := time.Now()

	filepaths := []string{}
	for _, path := range paths {
		err = filepath.Walk(path, func(fp string, f os.FileInfo, err error) error {
			if checkFiletype(fp) {
				filepaths = append(filepaths, fp)
				fmt.Println(fp)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("filepath.Walk error %v\n", err)
		}
	}

	for _, path := range filepaths {
		go convert(path)
	}

	fmt.Printf("total file count : %d\n", len(filepaths))
	fmt.Printf("time : %v\n", time.Since(start))

	time.Sleep(1 * time.Second)
}

func checkFiletype(file string) bool {
	for _, v := range filetype {
		if strings.HasSuffix(file, v) {
			return true
		}
	}
	return false
}

func convert(fp string) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range replaceText {
		data = bytes.Replace(data, []byte(k), []byte(v), -1)
	}

	err = ioutil.WriteFile(fp, data, 0666)
	if err != nil {
		fmt.Println(err)
	}

	command := dir + "/opencc-1.0.4/bin/opencc.exe -i " + fp + " -o " + fp + " -c " + dir + "/opencc-1.0.4/share/opencc/" + config + ".json"
	// cmd := exec.Command("cmd.exe", "/c", "start "+command)
	list := strings.Split(command, " ")
	cmd := exec.Command(list[0], list[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
