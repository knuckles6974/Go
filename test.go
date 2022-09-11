package main

// func square(x int) int {
// 	return 81
// }

// func main() {
// 	fmt.Printf("9 * 9 = %d\n", square(9))
// }
func nextValue() func() int {
	i := 0
	return func() int{
		i++
		return i
	}
}

func main() {
	next := nextValue()

	println(next())
	println(next())
	println(next())

	anoherNext := nextValue()
	println(anoherNext())
	println(anoherNext())
}