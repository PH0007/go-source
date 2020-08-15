package main

import (
	"fmt"
	"strconv"
)

//根据struct 拼接slice

type pp struct {
	x int
	y int
}

func (p1 *pp) getV() []string {
	var s [2]string

	s[0] = strconv.Itoa(p1.x)
	s[1] = strconv.Itoa(p1.y)
	p1.x = 100
	p1.y = 100
	s1 := s[:]
	return s1
}

func main() {
	var a1 pp
	a1.x = 30
	a1.y = 20

	val := a1.getV()
	fmt.Println(val)

	fmt.Println(a1)
}
