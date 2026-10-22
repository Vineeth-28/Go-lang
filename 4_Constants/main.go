package main

import "fmt"

const age = 12

func main() {

	const (
		port=5000
		host="localhost"
	)
	fmt.Println(age);
	fmt.Println(port,host);
}

// you can declare any varibles out side the function but not in case of declaring with shorthand ! you can declare multiple const at same time !