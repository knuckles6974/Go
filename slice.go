package main

import "fmt"

// func main() {
// 	var slice = []int{1,2,3}

// 	// for i := 0; i <len(slice); i++{
// 	// 	slice[i] += 10
// 	// }
// 	// fmt.Println(slice)

// 	// for i,v := range slice{
// 	// 	slice[i] = v *2
// 	// }
// 	// fmt.Println(slice)
// 	slice2 := append(slice, 4)

// 	fmt.Println(slice2)
// }

// func main(){
// 	var slice []int

// 	for i := 1; i <= 10; i++{
// 		slice = append(slice , i)
// 	}

// 	slice = append(slice, 11,12)

// 	fmt.Println(slice)
// }

func changearray(array2 [5]int) {
	array2[2]= 200

}

func changeslice(slice2 []int){
	slice2[2]= 200
}

func main(){
	array := [5]int{1,2,3,4}
	slice := []int{1,2,3,4}

	changearray(array)

	changeslice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}


