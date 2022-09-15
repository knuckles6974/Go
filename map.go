package main

import "fmt"


func main() {
	//맵: 해시테이블, 딕셔너리, key-value
	//래퍼런스 타입
	//비교연산자 사용 불가능 
	//특징 : 참조타입으로 사용 불가능, 값으로 모든타입 사용가능
	//make 함수 및 축약으로 초기화 가능
	//순서없음

	var map1 map[string] int = make(map[string] int) //정석적인 방법 
	var map2 = make(map[string]int) //자료형 생략
	map3 := make(map[string]int) //리터럴 형

	fmt.Println("exe1 :" , map1)
	fmt.Println("exe1 :" , map2)
	fmt.Println("exe1 :" , map3)
	fmt.Println()


	map4 := map[string]int{}
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple" :15,
		"banana" : 40,
		"orange" : 23,	
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println("ex2 :", map4)
	fmt.Println("ex2 :", map5)
	fmt.Println("ex2 :", map6)
	fmt.Println("ex2 :", map6["orange"])
	fmt.Println("ex2 :", map6["apple"])


}