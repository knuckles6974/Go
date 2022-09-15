package main

import "fmt"


func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot

}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2:" , value) 
	}
}



func main() {
	//함수 고급
	//가변 인자 실습

	x := multiply(5,6,7,8,9,3)

	fmt.Println("ex1:", x)

	prtWord("a", "apple", "test", "golang","seoul")

	//예제

	a := []int{5,6,7,8,9,10}

	m := multiply(a...)
	fmt.Println("ex3 :", m)

	
}
