package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type proc struct {
	task     string
	PID      string
	taskType string
	cap      string
	mem      string
	unit     string
}

func main() {
	cmd := exec.Command("tasklist")

	// err := cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// 因为结果是字节数组，需要转换成string
	// fmt.Println(string(output))
	re, _ := regexp.Compile(" +")
	newStr := re.ReplaceAllString(string(output), " ")

	fmt.Println(newStr)
	s := strings.Split(string(newStr), "\n")

	var pros []proc
	var pro1 proc
	for _, v := range s {

		var tmp = strings.Split(v, " ")
		if len(tmp) > 4 {
			pro1.task = tmp[0]
			pro1.PID = tmp[1]
			pro1.taskType = tmp[2]
			pro1.cap = tmp[3]

			pros = append(pros, pro1)
		}

	}
	fmt.Println(pros)
}
