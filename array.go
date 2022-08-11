// package main

// import "fmt"

// func main(){
// 	var t [5]float64 =  [5]float64{24.0,25.9, 27,8,26.9, 26.2}

// 	for i , v := range t{
// 		fmt.Println(i,v)
// 	}
// }
package main

import "fmt"


func main(){
	a := [2][5]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
	}
	for _, arr := range a {
		for _, v := range arr{
			fmt.Print(v,"")
		}
		fmt.Println()
	}
}
