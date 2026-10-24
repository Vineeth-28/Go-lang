package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
//if the parameter are same we can write once (a,b int), in function we can return multiple value in go
func getlanguages() (string, string, string) {
	return "golangauage", "javascript", "typescript"

}
//should be of same order in multiple return ,if we dont have to use any values we put _

func main() {
	fmt.Println("hello")
	fmt.Println(add(2, 3))
	lang1, lang2, lang3 := getlanguages()
	fmt.Println(lang1, lang2, lang3)

}
