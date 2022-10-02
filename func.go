package main

import "fmt"



func addTo(base int , vals ...int) []int{
	out := make([]int,0, len(vals))
	for _, v := range vals{
		out = append(out, base+v)
	}
	return out
}

func main(){
	fmt.Println(addTo(3))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))	
}