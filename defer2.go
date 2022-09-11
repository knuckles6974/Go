package main

import "fmt"

// func sayHello(msg string) {
// 	defer func(){
// 		fmt.Println(msg)
// 	}()

// 	func() {
// 		fmt.Println("Hi")
// 	}()
// }

// func main() {
// 	//예제1
// 	sayHello("Golang")
// }
func stack() {
	for i := 1; i <=10; i++{
		defer fmt.Println("ex1: ",i)
	}
}



func main() {
	
	stack()
}