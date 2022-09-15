package main

import (
	"fmt"
	"strconv"
)

func sum(x int, y int) int{
	return x + y
}



func main() {
	//함수
	//func 함수명(매개변수) (반환타입or 반환값 변수명) : 반환값존재
	//func 함수명() (반환타입or 반환값 변수명) : 반환값존재
	//func 함수명(매개변수) : 매개변수 존재, 반환겂 없음
	//func 함수명() : 매개변수 없음, 반환값 없음
	//타 언어와 달리 반환 값 다수 가능


	//예제
	result := sum(5,5)
	fmt.Println("ex3:", result)
	fmt.Println("ex3:", strconv.Itoa(sum(5,5)))

	//예제

	// x1,x2,x3,x4,x5 := arrayMultiply(1,2,3,4,5)
	// y1, _,y3,_, y5 := arrayMultiply(1,2,3,4,5)
	// fmt.Println("ex2:", x1,x2,x3,x4,x5)
	// fmt.Println("ex2:", y1,x2,x3,x4,x5)




}

