package main

import "fmt"


type Dog struct {
	name string
	weight int
}
//예제1
func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

type Behavior interface {
	bite()
}

func main() {
	dog1 := Dog{"poll",10}

	var inter1 Behavior
	inter1 = dog1
	inter1.bite()

	//예제2
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)

	inter2.bite()

	//예제3
	inters := []Behavior{dog1, dog2}

	//인덱스 형태로 실행
	for idx, _ := range inters {
		inters[idx].bite()
	}

	// 값 형태로 실행(인터페이스)
	for _, val := range inters {
		val.bite()
	}

}


