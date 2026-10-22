package main

import "fmt"

func main() {
	var nums [4]int

   nums[0]=1
   //array length 
	fmt.Println(len(nums))
   fmt.Println(nums[0])
   fmt.Println(nums)

   var names[2]string 

   names[0]="john"
   names[1]="doe"
   fmt.Println(names);

   numes:= [3]int {1,23,4}; //Declare in single line
   fmt.Println(numes); 

   //2d Arrays
   number:= [2][2]int {{4,2},{5,6}}; 
   fmt.Println(number)

   //we mostly use slices instead of arrays to optimize , fixed size we know then its array 
   

 }
