package main

import (
	"fmt"
	"io"
	"os"
)

// 把文件copy到另外一个文件夹
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	//拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	fmt.Println("开始复制...")
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("报错了!")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}

func main() {
	var fileFrom string = "test1.csv"
	var fileTo string = "./lint/test2.csv"

	a, _ := CopyFile1(fileFrom, fileTo)
	fmt.Println(a)
}
