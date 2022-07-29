package main

import "fmt"

// func HasRichFriend() bool{
// 	return true
// }

// func GetFriendsCount() int{
// 	return 3
// }

// func main() {
// 	price := 35000

// 	if price > 50000 {
// 		if HasRichFriend() {
// 			fmt.Println("신발끈")
// 		} else {
// 			fmt.Println("앤빵하자")
// 		}
// 	} else if price >= 30000 && price <= 50000 {
// 		if GetFriendsCount() > 3{
// 			fmt.Println("신발끈이")
// 		} else {
// 			fmt.Println("나눠내자")
// 		}

// 		} else {
// 			fmt.Println("오늘은 내가쏜다")
// 		}
// }

// func getMyAge() (int, bool) {
// 	return 33, true
// }

// func main () {

// 	if age, ok := getMyAge(); ok && age < 20{
// 		fmt.Println("Young", age)

// 	} else if ok && age < 30{
// 		fmt.Println("Nice Age", age)
// 	} else if ok {
// 		fmt.Println("You are ugly" , age)
// 	} else {
// 		fmt.Println("Error")
// 	}
// 	fmt.Println("your age is", age)
// }


func main(){
	temp := 30
	rain := 40

	if temp >= 25 && rain >= 80{
		fmt.Println("비옴")
	} else if rain >= 20{
		fmt.Println("덥고습")
	} else if temp >=25 && rain >= 20 {
		fmt.Println("야외활동굿")
	} else if temp < 10 || rain >= 80 {
		fmt.Println("야외활동 노")
	} else {
		fmt.Println("좋은날씨")
	}
}