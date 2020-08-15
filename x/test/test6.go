package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := ColoredPoint{Point{2, 3}, 2}
	Struct2Map(a)

}

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color int
}

func Struct2Map(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	fmt.Println("----------------")
	fmt.Println(t)
	fmt.Println(v)
}
