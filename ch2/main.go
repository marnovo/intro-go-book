// Chapter 2. Types

package main

import "fmt"

func main()  {
	// Strings
	fmt.Println(len("Hello, World"))
    fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")
	
	// Booleans
	fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
