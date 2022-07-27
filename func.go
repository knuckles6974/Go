package main

import "fmt"

// func Add(a int, b int) int{

// 	return a + b
// }

// func main() {
// 	c := Add(3, 6)
// 	fmt.Println(c)
// }

// func main(){
//     math := 80
// 	eng := 54
// 	history := 65
// 	fmt.Println("고퍼님의 평균점수는", (math + eng + history)/3, "입니다.")
// }

// func PrintAvgScore(name string, math int, eng int, history int) {
// 	total := math + eng + history
// 	avg := total / 3
// 	fmt.Println(name, "님 평균 점수는", avg, "입니다.")
// }

// func main(){
// 	PrintAvgScore("김일등", 80,54,65)
// 	PrintAvgScore("김일", 40,24,65)
// }

// func Divide(a, b int) (int, bool){
// 	if b == 0{
// 		return 0, false
// 	}
// 	return a / b, true
// }

// func main(){
// 	c, success := Divide(9,3)
// 	fmt.Println(c, success)
// 	d, success := Divide(9,0)
// 	fmt.Println(d, success)
// }

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return 
	}
	result = a / b
	success = true
	return 
}

func main(){
	c, success := Divide(9,3)
	fmt.Println(c, success)
	d, success := Divide(9,0)
	fmt.Println(d, success)
}