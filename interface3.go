package main

import "fmt"

type Dog struct{
	name string
	weight int
}

type Cat struct {
	name string
	weight int
}

//구조체 Dog 메소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, ": Dog bites!")
}
func (d Dog) sounds() {
	fmt.Println(d.name, ": Dog barks!")
}
func (d Dog) run() {
	fmt.Println(d.name, ": Dog running!")
}

//구조체 Cat 메소드 구현
func (d Cat) bite() {
	fmt.Println(d.name, ": Cat bites!")
}
func (d Cat) sounds() {
	fmt.Println(d.name, ": Cat criess!")
}
func (d Cat) run() {
	fmt.Println(d.name, ": Cat running!")
}

type Behavior interface {
	bite()
	sounds()
	run()
}

//인터페이스를 파라미터 받는다.
func act(animal Behavior){
	animal.bite()
	animal.sounds()
	animal.run()
}


func main() {
	//예쩨1
	dog := Dog{"poll", 10}
	cat := Cat{"bob",5}

	//개 행동 실행
	act(dog)

	//개 행동 실행
	act(cat)
}