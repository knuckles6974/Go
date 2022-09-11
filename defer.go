package main

import "fmt"

// func ex_f1() {
// 	fmt.Println("f1 : start")
// 	defer ex_f2()
// 	fmt.Println("f1: end")
// }

// func ex_f2() {
// 	fmt.Println("f2: called")
// }

// func main() {
// 	//defer 함수 실행
// 	//defer를 호출한 함수가 종료되기 직전에 호출된다.
// 	//타 언어의 finally문과 비슷
// 	//주로 리소스 반환, 열린 파일 닫기, mutex잠금 해제

// 	//예제1

// 	ex_f1()
// }


func start(t string) string{
	fmt.Println("start:", t)
	return t
}

func end(t string) {
	fmt.Println("end :", t)
}


func a() {
	defer end(start("b"))
	fmt.Println("in a")
}


func main() {
	//예제1
	a()
}