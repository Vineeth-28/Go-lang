package main

import "fmt"

func main() {
	nums := []int{1, 5, 14, 1}
	//iterating with for loop

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	//range keyword
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println(sum)
	// range keyword with index i
	for i, num := range nums {
		fmt.Println(num, i)
	}

	//iterating over map
	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m { // for convention in go we use name  k and v
		fmt.Println(k, v)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
