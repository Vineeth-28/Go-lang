package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num 
	}
	return total
}

func main() {
	//nums:=[]int[3,13,1,31]; 
    // result:=sun[nums...]   
	result := sum(1, 3, 1, 3, 1, 4)
	fmt.Println(result) 
}
