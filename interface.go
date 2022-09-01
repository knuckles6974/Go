package main

import "fmt"

// type DuckInterface interface {
// 	Fly()
// 	Walk(distance int) int
// }

// type Sample interface {
// 		String() string
// 		String(int) string
// 		_(x int)
// }

// type Stringer interface {
// 	String() string
// }

// type Student struct {
// 	Name string
// 	Age  int

// }
// func (s Student) String() string {
// 	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age,s.Name)
// }

// func main(){
// 	student := Student{"철수", 12}
// 	var stringer Stringer

// 	stringer = student

// 	fmt.Printf("%s\n", stringer.String())
//

// func PrintVal(v interface{}) {
// 	swith t := v.(type) {
// 	case int :
// 		fmt.Printf("v is int %d\n", int(t))
// 	case float64:
// 		fmt.Printf("v is float %f\n", float64(t))
// 	case string:
// 		fmt.Printf("v is string %s\n", string(t))
// 	default:
// 		fmt.Printf("Not supported type: %T:%v\n", t,t)
// 	}
// }
// type Student struct {
// 	Age int
// }

// func main(){
// 	PrintVal(10)
// 	PrintVal(3.14)
// 	PrintVal("Hello")
// 	PrintVal(Student{15})
// }

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {

	s := stringer.(*Student)
	fmt.Printf("Age:%d\n", s.Age)
}

func main() {
	s := &Student{15}

	PrintAge(s)
}




