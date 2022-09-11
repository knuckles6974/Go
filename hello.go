package main

import "fmt"

func main(){
	a := argTest(3, 3)
	fmt.Println(a)
}

func argTest(arg , argA int32) int32 {
	result :=  arg + argA
	return result
}