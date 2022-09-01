package main

import (
	"fmt"
	"time"
)


func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}
func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}

func main(){
	//고루틴
	//타 언어의 쓰레드와 비슷한 기능
	//생성 방법 매우 간단, 리소스 매우 적게 사용
	//즉, 수맣은 고루틴 동시 생성 실행 가능
	//비동기적 함수 루틴 실행() -> 채널을 통한 통신 가능
	//공유메모리 사용시에 정확한 동기화 코딩 필요

	exe1() //가장먼저 실행
	

	fmt.Println("Main Routine Start", time.Now())
	go exe2()
	time.Sleep(4 * time.Second)
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine Start", time.Now())
}