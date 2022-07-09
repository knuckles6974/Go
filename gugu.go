package main

import "fmt"

func main(){
	for i := 2 ; i < 10 ; i++{
		for j := 2 ; j<10 ; j++{
			fmt.Printf("%d * %d =%d\t" , j,i,i*j)
		}
	}
}