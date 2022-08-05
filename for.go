package main


import "fmt"

// func main(){
// 	for i := 0; i < 5; i++{
// 		for j := 0; j < i+1; j++{
// 			fmt.Println("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func main(){
// 	dan := 2
// 	b:= 1
// 	for {
// 	for {
// 		fmt.Printf("%d * %d = %d\n", dan , b, dan*b)
// 		b++
// 		if b == 10 {
// 			break
// 		}
// 	}
// 	b = 1
// 	dan++
// 	if dan == 10{
// 		break
// 	}
// 	}
// 	fmt.Println("for문이 종료")
// }

// func main(){
// 	a := 1
// 	b := 1

// OuterFor:
// 	for ; a <=9; a++{
// 		for b = 1; b<=9; b++{
// 			if a * b == 45{
// 				break OuterFor
// 			}
// 		}
		
// 	}
// 	fmt.Printf("%d * %d = %d\n",a,b, a+b)
// }

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++{
		if a * b ==45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0

	for ; a <= 9; a++{
		var found bool
		if b, found = find45(a); found{
			break
		}
	}
	fmt.Printf("%d* %d= %d\n",a,b,a*b)
}