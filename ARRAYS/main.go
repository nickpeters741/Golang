package main

import "fmt"

//arrays
func todo() {
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"am an new"}

	arr2 = append(arr2, "gopher")
	fmt.Println(arr, arr2)
}

func main() {
	todo()
	//fmt.Println()
}
