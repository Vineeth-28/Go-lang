package main

import "fmt"

func main() {
   //simple switch  no need to mention break 
   i:=5

   switch i{
   case 1: fmt.Println("one");
   case 2: fmt.Println("two");
   case 3: fmt.Println("three");
   default: fmt.Println("not one two three")
   }

}
