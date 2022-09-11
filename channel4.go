package main

import (
	"fmt"
	"runtime"
)

// func rangeSum(rg int, c chan int) {
// 	sum := 0

// 	for i := 0; i <= 1000; i++{
// 		sum += i
// 	}
// 	c <- sum
// }

// func main() {

// 	c := make(chan int)

// 	go rangeSum(1000, c)
// 	go rangeSum(700, c)
// 	go rangeSum(500,c)

// 	result1 := <- c
// 	result2 := <- c
// 	result3 := <- c
// 	fmt.Println("ex1 :", result1)
// 	fmt.Println("ex1 :", result2)
// 	fmt.Println("ex1 :", result3)
// }

// func main() {

// 	ch := make(chan bool)
// 	cnt := 6

// 	go func() {
// 		for i := 0; i <cnt; i++{
// 			ch <- true
// 			fmt.Println("Go:" ,i)
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	for i := 0;i <cnt; i++{
// 		<-ch
// 		fmt.Println("main:" ,i)
// 	}
// }

func main() {

	runtime.GOMAXPROCS(1)
	ch := make(chan bool,2)
	cnt := 12

	go func() {
		for i := 0; i <cnt; i++{
			ch <- true
			fmt.Println("Go: ",i)
		}
	}()
	for i := 0; i <cnt; i++{
		<-ch
		fmt.Println("Main: ",i)
	}
	
}