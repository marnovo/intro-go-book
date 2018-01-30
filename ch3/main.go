// Chapter 3. Variables

package main

import "fmt"

func main()  {
	// working with variables
	var x string
	x = "first "
	x = x + "second"
	fmt.Println(x)

	// Calling further functions
	multiVar()
	doubleIt()
}

func multiVar()  {
	
	// Constants
	const msg string = "Hello, World"

	// Declaring multiple variables
	var (
		a = 5
		b = 10
		c = 15
	)
	
	fmt.Println(msg)
	fmt.Println(a + b + c)

}

func doubleIt() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input * 2

    fmt.Println(output)
}
