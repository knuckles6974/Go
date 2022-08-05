package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil{
			fmt.Println("숫자입력")

			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신숫자는 %d 입니다.\n",number)
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료.")
}