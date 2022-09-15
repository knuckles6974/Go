package main

import "fmt"

func multiply(x int, y int) (r1 int, r2 int) { 
	r1 = x * 10
	r2 = y * 20

	return r1 , r2
}

// func arrayMultiply(a,b,c,d,e int) (int, int,int,int,int) {
// 	return a * 1 , b * 2, c *3 , e * 5, d * 5 
// }

func main() {
	//다중값 반환


	//에제

	a, b := multiply(10, 5)
	
	fmt.Println("ex : " ,a, b)
}