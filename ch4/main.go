// Chapter 4. Control Structures

package main

import "fmt"

func main(){
	// Loops

	// Simple loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// More compact version
	for i := 1; i <= 10; i++ {
        fmt.Println(i)
	}
	
	// If statements
	x := 3
	if x == 1 {
		fmt.Println("1")
	} else if x == 2 {
		fmt.Println("2")
	} else {
		fmt.Println("3 or other")
	}

	// Loop + if statements
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	// Switch statement
	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
	}

	/*
	Ex 3
	Write a program that prints the numbers from 1 to 100, but for multiples 
	of three, print “Fizz” instead of the number, and for the multiples of 
	five, print “Buzz.” For numbers that are multiples of both three and five, 
	print “FizzBuzz.
	*/
	fmt.Println("\rExercise 3:\r")
	for k := 1; k <= 100; k++ {
		if (k % 3 == 0) && (k % 5 == 0) {
			fmt.Println("FizzBuzz")
		} else if (k % 3 == 0) {
			fmt.Println("Fizz")
		} else if (k % 5 == 0) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(k)
		}
	}

}
