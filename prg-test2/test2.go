package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s1 = "good mornign\ndddd\t fdafdafe"
	re, _ := regexp.Compile(" +")
	r := re.ReplaceAllString(s1, "-")
	fmt.Println(r)
}
