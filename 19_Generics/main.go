//Generics 

package main

import "fmt"

// T is a type parameter. 
// "int | float64" is the constraint limiting what T can be.
func Sum[T int | float64](a, b T) T {
    return a + b
}

func main() {
    // Explicit instantiation
    fmt.Println(Sum[int](5, 10)) 
    
    // Type Inference (Go automatically detects the types)
    fmt.Println(Sum(3.14, 2.71)) 
}


//Type and sets

package main

// Custom type constraint interface
type Numeric interface {
    ~int | ~int64 | ~float64
}

type ID int // Custom type with underlying type int

func PrintAbsolute[T Numeric](val T) {
    // Works with int, float64, and custom types like ID because of '~'
    if val < 0 {
        val = -val
    }
    println(val)
}
