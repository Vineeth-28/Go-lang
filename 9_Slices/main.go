package main

import "fmt"

func main() {

	//uninitialized slices is nil
	var nums []int
	fmt.Println(nums==nil)
    // Capcity - max number of element can fit  dyanmic arrays 
	var number = make([]int , 4 ,5)
    fmt.Println(number)
}
