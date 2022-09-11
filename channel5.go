package main

import (
	"fmt"
)


func main(){
	//close : 채널닫기, 주의 -> 닫힌 채널에 값 전송시 패닉 발생
	//range : 채널안에서 값을 꺼낸다. , 채널을 닫아야 반복문 종료 -> 채널이 열려있고 값 전송하지 않으면 계속대기
	ch := make(chan bool)
	go func(){
		for i:= 0; i<5; i++{
			ch <- true

		}
		close(ch)
	}()
	for i := range ch{ //채널에서 값을 꺼내온다.채널에서 close될뗴까지
		fmt.Println("ex1:",i)
	}
}