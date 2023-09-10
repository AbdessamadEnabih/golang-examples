package loops

import "fmt"

func main() {
	// Three-component loop
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	// switch
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	// foreach loop
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	for {
		fmt.Println("loop")
		break
	}
}
