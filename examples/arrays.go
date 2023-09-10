package array

import "fmt"

func Arrays() {
	var arr1 = [3]int{1, 2, 3}      // definded length
	arr2 := [...]int{4, 5, 6, 7, 8} // inferred length

	fmt.Println(arr1)
	fmt.Println(arr2)
}
