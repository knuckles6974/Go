package main

import "fmt"

// func main(){
// 	str := "Hello 월드!"
// 	arr := []rune(str)

// 	for i := 0; i < len(str); i++{

// 		fmt.Printf("타입:%T 값:%d 문자값:%c\n" , str[i],  str[i], str[i])

// 	}
// }

// func main(){
// 	str := "Hello 월드!"
// 	arr := []rune(str)

// 	for i := 0; i < len(arr); i++{
// 		fmt.Printf("타입:%T 값:%d 문자값:%c\n" , arr[i],  arr[i], arr[i])
// 	}
// }


func main(){
	str := "Hello 월드!"
	for _, v := range str {
		fmt.Printf("타입:%T 값:%d 문자:%c\n",v,v,v)
	}
}