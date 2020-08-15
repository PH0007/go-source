package main

import (
	"fmt"
	"reflect"
)

type s1 struct {
	x int
	y int
	z int

	sum int
}

func main() {
	p := s1{2, 3, 4, 32}

	// l: = len(p)
	s := reflect.TypeOf(p).Elem()
	// for i := range s {

	fmt.Println(s)

	// }

}
