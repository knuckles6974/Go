package main

import "fmt"


func main() {
	//맵 조회 할 경우에 주의 할 점

	//예제1
	map1 := map[string]int{
		"apple" :15,
		"banana" :115,
		"orange" : 1115,
		"lemon" :0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println("ex1 :", value1)
	fmt.Println("ex1 :", value2)
	fmt.Println("ex1 :", value3,ok)


	if value , ok := map1["kiwi"]; ok {
		fmt.Println("ex2: ",value)
	} else{
		fmt.Println("ex2:, kiwi is not exist!")
	}	
	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2:, kiwi is not exist!")
	}
}