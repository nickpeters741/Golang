package main

import (
	"fmt"
	"strings"
)

func main() {
	var m1 string
	var m2 string
	m1 = "my name"
	m2 = "my"
	fmt.Println(strings.Contains(m1,m2))
}
