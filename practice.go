// package main

// import "fmt"


// func main(){
	

// 	for a := 10; a > 0; a--{
// 	fmt.Print(a," ")
// }
// }

package main

import "fmt"

func main(){
	for i := 0; i < 5; i++{
		for j := 0; j<5-i; j++{
			fmt.Println("*")
		}
	}
}
