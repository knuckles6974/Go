package main

import "fmt"

func main(){
	var a uint8
	a |= 2
	a |= 4
	a |= 8

	var b uint8
	b = 4
	
	
	fmt.Printf("숫자 : %d\n 숫자2: %d\n", a,b)
}