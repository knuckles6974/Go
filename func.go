package main

import "fmt"


func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a int, b int) {
	fmt.Println("ex1:",a+b)
}

func multi_value(i int) {
	i = i *10
}


func multi_reference(i *int){
	*i  *= 10 //
}


func main() {
	//함수 전달, 참조전달, 값전달
	//예제1

	sum(100,add)//매개변수로 함수 전달
	//예제2

	a := 100
	multi_value(a)
	fmt.Println("ex2:",a)

	//예제3
	b := 100
	multi_reference(&b)
	fmt.Println("ex3:", b)



}