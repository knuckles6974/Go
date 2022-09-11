package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//뮤텍스 : 상호 배제 -> Thread들이 서로 running time에 서로 영향을 주지 않게 즉, 단독으로 실행되게 하는 기술
	//뮤텍스 : 여러 고루틴에서 작업하는공유 데이터 보호 
	//RWMutex : 쓰기 lock - 쓰기 시도중에는 다른 곳에서 이전 값을 읽으면 x, 읽기 락, 쓰기 락 전부 방지
	//RMutex  ;읽기 lock -> 읽기 시도 중에 값이 변경 방지 즉, 쓰기 락 방어 

	//동기화 사용하지 않은 경우 예제
	//쓰기 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환할 가능서 ㅇ증가
	//시스템 전체 cpu사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)
	
	go func() {
		for i := 0; i <=10; i++{
			data += 1
			fmt.Println("Write:", data)
			time.Sleep(200* time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <=10; i++{
			data += 1
			fmt.Println("Read1:", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <=10; i++{
			data += 1
			fmt.Println("Read2:", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}