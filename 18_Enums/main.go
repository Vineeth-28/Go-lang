package main

import "fmt"

//Define a custom type
type Status int 

const (
	statusPending Status = iota //0
	StatusActive 
	StatusFailed
	StatusUnknown

)

func main(){
	var current Status =StatusActive
  if(current ==StatusActive){
	fmt.Println("The account is active ")
  }	
}
