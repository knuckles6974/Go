package main

import "fmt"

// type Cloth struct{
// 	Tee string
// 	Pants string
// 	Total_price int
// }

// func main(){
// 	var cloth Cloth
// 		cloth.Tee = "margiela"
// 		cloth.Pants = "levis"
// 		cloth.Total_price = 5000000

// 	fmt.Println("티샤쓰:" , cloth.Tee)
// 	fmt.Println("바지:", cloth.Pants)
// 	fmt.Println("합계:", cloth.Total_price)
// }

type User struct{
	Name string
	ID   string
	Age  int
}
type VIPUser struct{
	UserInfo User
	VIPLevel int
	Price    int
}

func main(){
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑","hwarang", 40},
		3,
		250,
	}
	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID,user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이: %d VIP레벨: %d VIP가격: %d 만 원\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,

	)	
}


