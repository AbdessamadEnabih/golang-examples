package variable

import "fmt"

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	var a string
	var b int
	var c bool

	var a, b, c, d int = 1, 3, 5, 7 // multiple decl

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// const
	const A int = 1

	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

}
