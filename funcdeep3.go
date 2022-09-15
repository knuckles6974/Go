package main

import "fmt"

// func fact(n int) int{
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * fact(n-1)
// }

// func prtHEllo (n int) {
// 	if n == 0 {
// 	   return
// 	}
// 	fmt.Println("ex2:  (",n,")","hi")
// 	prtHEllo(n-1)
// }


func main() {
	//함수 고급
	//재귀함수
	//익명함수
	//즉시실행

	//예제1
	func() {
		fmt.Println("ex1: an")
	}()


	//예제2
	 msg := "hello golang"

	 func(m string) {
	 fmt.Println("ex2 :" , m)
	}(msg)

	func(x, y int) {
		fmt.Println("ex3: ", x + y)
	}(10, 20)

	//예제3
	r := func(x, y int) int{
		return x * y
	}(10, 100)
	
	fmt.Println("ex4:", r)
}

	
	
	//prtHEllo(10)

