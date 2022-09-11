package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main(){
// 	stdin := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Println("입력")
// 		var number int
// 		_, err := fmt.Scanln(&number)
// 		if err != nil{
// 			fmt.Println("숫자입력")

// 			stdin.ReadString('\n')
// 			continue
// 		}
// 		fmt.Printf("입력하신숫자는 %d 입니다.\n",number)
// 		if number%2 == 0 {
// 			break
// 		}
// 	}
// 	fmt.Println("for문이 종료.")
// }

// import "fmt"

// func main(){
// 	var str string = "Hello World"
// 	var slice []byte = []byte(str)

// 	slice[2] = 'a'

// 	fmt.Println(str)
// 	fmt.Printf("%s\n",slice)
// }

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main(){
// 	rand.Seed(time.Now().UnixNano())

// 	n := rand.Intn(100)
// 	fmt.Println(n)

// }

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker
	att.Attack()
}