package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){

	//클로저사용예쩨
	//예제1
	runtime.GOMAXPROCS(1)

	s := "Goroutune Closure :"

	for i := 0; i < 1000; i++{
	 go	func(n int) {
			fmt.Println(s, n, "-", time.Now())
		}(i)
	}

	time.Sleep(5 * time.Second)
}