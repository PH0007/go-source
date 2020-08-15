package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type person struct {
	id   int
	name string
	age  int
}

func (per1 person) per2S() [3]string {
	s1 := [3]string{}

	for k, v := range per1 {
		s1[k] = v
	}
	return s1
}

func main() {

	p1 := []person{{10, "jim", 32}, {20, "lili", 21}, {23, "jfja", 3}}
	p2 = [3][3]string{}

	for i, _ := range p1 {
		p2[i] = p1[i].per2S()
	}
	f, err := os.Create("test.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流
	data := p2
	fmt.Println(data)
	w.WriteAll(data) //写入数据
	w.Flush()

	// file, err := os.Open("test1.csv")
	// if err != nil {
	//     fmt.Println("Error:", err)
	//     return
	// }
	// defer file.Close()
	// reader := csv.NewReader(file)
	// for {
	//     record, err := reader.Read()
	//     if err == io.EOF {
	//         break
	//     } else if err != nil {
	//         fmt.Println("Error:", err)
	//         return
	//     }
	//     fmt.Println(record) // record has the type []string
	// }
}
