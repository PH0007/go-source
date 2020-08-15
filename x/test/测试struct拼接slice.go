package main

//根据struct 拼接slice
import (
	"fmt"
)

type point struct {
	x int
	y int
}

func main() {
	a1 := []point{{2, 3}, {2, 4}, {2, 5}}
	fmt.Println(a1)

	a1 = append(a1, point{2, 8})

	for i := range a1 {
		fmt.Println(a1[i].y)

	}

}
