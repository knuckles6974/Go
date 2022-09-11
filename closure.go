package main

import "fmt"

//클로저 는 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
//함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언 된 변수에 접근 가능한 함수
//함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷

//예제1

// 	multifly := func(x, y int) int {
// 		return x * y
// 	}

// 	r1 := multifly(5,10)

// 	fmt.Println("ex1 : ", r1)

// 	m, n := 5,10
// 	sum := func(c int) int{
// 		return m + n + c
// 	}

// 	r2 := sum(10)
// 	fmt.Println("ex2 :", r2)

// }


func main() {
	cnt := increaseCnt()

	fmt.Println("ex1:", cnt())
}

func increaseCnt() func() int {
	n := 0 //지역변수
	return func() int{
		n += 1
		return n 
	}
}