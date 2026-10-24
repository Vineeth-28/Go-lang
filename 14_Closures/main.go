package main

import "fmt"

func counter() func() int {
    var count int = 0
    return func() int {
        count += 1 
        return count
    }
}

func main() {
    increament := counter()
    fmt.Println(increament()) // Outputs: 1
    fmt.Println(increament()) // Outputs: 2
    fmt.Println(increament()) // Outputs: 3
}
