package main

import "fmt"
func main(){
	M1 := 2
	ptr :=&M1
	//ptr simply prints address however when you add *,it prints the value stored in it
	fmt.Println(*ptr)
}