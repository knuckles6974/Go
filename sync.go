package main

import (
	"fmt"
	"runtime"
	"sync"
)

type count struct{
	num int
	mutex sync.Mutex
}


func (c *count) increment(){
	//공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num += 1
	//공유 데이터 수정 후 보호 해제
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println(c.num)
}
func main() {
	//실행흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장중요
	//뮤텍스(mutex) :여러 고루틴에서 작업하는 공유데잍터 보호
	//sync.Mutex 선언 후 lock, Unlock 사용
	//동기화 사요하지 않은 경우 예제
	//시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num :0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++{
		go func(){
			c.increment()
			done <- true
			runtime.Gosched() //cpu 양보

		}()
	}

	for i := 1; i <= 10000; i++{
		<-done
	}

	c.result()
}