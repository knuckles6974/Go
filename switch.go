package main

import (
	"fmt"
	"math/rand"
)

// type ColorType int
// const (
// 	Red ColorType = iota
// 	Blue
// 	Green
// 	Yellow
// )

// func colorToString(color ColorType) string {
// 	switch color {
// 	case Red:
// 		return "Red"
// 	case Blue:
// 		return "Blue"
// 	case Green:
// 		return "Green"
// 	case Yellow:
// 		return "Yellow"
// 	default:
// 		return "Undefined"
// 	}
// }

// func getMyFavoriteColor() ColorType {
// 	return Blue
// }

// func main() {
// 	fmt.Println("my fav color is" , colorToString(getMyFavoriteColor()))
// }


func main() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 ==0{
			goto done
		}
		a = a*2 +1
	}
	fmt.Println("do")
done:
	fmt.Println("efe")	
}	