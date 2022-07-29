package main

import "fmt"

// func main() {
// // 	a:=3

// // 	switch a {

// // 	case 1:
// // 		fmt.Println("a == 1")
// // 	case 2:
// // 		fmt.Println("a == 2")
// // 	case 3:
// // 		fmt.Println("a == 3")
// // 	case 4:
// // 		fmt.Println("a == 4")
// // 	default:
// // 		fmt.Println("a > 4")
// // 	}
// }

// func getMyAge() int {
// 	return 22
// }

// func main() {
// 	switch age := getMyAge(); {
// 	case age < 10:
// 		fmt.Println("child")
// 	case age < 20:
// 		fmt.Println("teenager")
// 	case age < 30:
// 		fmt.Println("20s")
// 	default:
// 		fmt.Println("My age is", age)
// 	}

// }

type ColorType int
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"			 		
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("my fav color is" , colorToString(getMyFavoriteColor()))
}

