package main

import "fmt"
	   "time"

// type account struct{
// 	balance int
// }

// func withdrawFunc(a *account, amount int) {
// 	a.balance -= amount
// }

// func (a *account) withdrawMethod(amount int){ //1번
// 	a.balance -= amount
// }

// func main(){
// 	a := &account{ 100 }

// 	withdrawFunc(a, 30)

// 	a.withdrawMethod(30)

// 	fmt.Printf("%d\n", a.balance)

// }

// type myInt int

// func (a myInt) add(b int) int{
// 	return int(a)+ b
// }

// func main(){
// 	var a myInt = 10
// 	fmt.Println(a.add(30))
// 	var b int = 20
// 	fmt.Println(myInt(b).add(50))
// }

type Courier struct{
	Name string
}

type Product struct{
	NAme string
	Price int
	ID int
}

type Parcel struct{
	Pdt *Product
	ShippedTime time.Time
	DeliveredTime time.Time

}

func (c *Courier) SendProduct(pdt *Product) *Parcel{
	p := &Parcel{}
	p.Pdt = pdt
	p.ShippedTime = time.Now()
	return p

}

func (p *Parcel) Delivered() *Product{
	p.DeliveredTime = time.Now()
	return p.Pdt
}




