/*
	This is my 1st go project

File: hello.go
Description: This program prints "Hello World" to the console and that's just the start
Author: Gatheru Oscar
Date: 2025-07-13
*/

package main

import (
	"errors"
	"fmt"
)

var name string = " from me!"
var hour int = 1

func panicExample() {
	fmt.Println("Who is this?")
	defer fmt.Println("ooh! This is not the owner")
	//panic("Something went wrong")
}

func example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	panic("Something went soo wrong")
}

/*
	func divide(a, b int) int {
		if b == 0 {
			defer fmt.Println("Who taught you in school?")
			//panic("Division by zero")
		}
		return a / b
	}
*/

// dealing with functions
func add(a, b int) int { //basic function
	return a + b
}

// multiple parameters and types
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// named return types
func areaAndPerimeter(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

// Variadic function(basic)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// variadic function with other parameters
func printMessage(prefix string, messages ...string) {
	for _, msg := range messages {
		fmt.Println(prefix, msg)
	}
}

// Multiple return values
func divides(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Handling multiple return types
func getCoordinates() (int, int, string) {
	return 10, 20, "North"
}

// Closure
func counter() func() int { //scope of the anonymous fn
	count1 := 0
	return func() int {
		count1++ //accessess the variable declared in the outer counter() scope
		return count1
	}
}

// Recursion
func factorial(a1 int) int {
	if a1 == 0 { //base case
		return 1
	}
	// Recursive step
	return a1 * factorial(a1-1) //Basically a1! = a1*(a1-1)
}
func fib(a2 int) int {
	if a2 <= 1 {
		return a2
	}
	return fib(a2-1) + fib(a2-2) //Basically Fib(a2) = fib(a2-1) + fib(a2-2)
}
func main() {
	k, h, _ := getCoordinates()
	fmt.Println("K: ", k, "H: ", h)
	_, _, direction := getCoordinates()
	fmt.Println("Direction: ", direction) //example usage of handling return types
	result := add(1, 1)
	q, r := divide(7, 5)
	ar, p := areaAndPerimeter(4, 5)
	quotients, errr := divides(13.0, 6.5)
	if errr != nil {
		fmt.Println("Error: ", errr)
	} else {
		fmt.Println("Quotient: ", quotients)
	}
	_, errr = divides(5.0, 0.0)
	if errr != nil {
		fmt.Println("Error: ", errr)
	}
	printMessage("Hi!", "Welcome to Golang", "Here we're working on the basics of the language", "Where are you from?", "I meant language")
	fmt.Println("Sum: ", sum(1, 3, 5, 8))
	fmt.Println("Area: ", ar, "Perimeter: ", p)
	fmt.Println("Quotient: ", q, "Remainder: ", r)
	fmt.Println("Result: ", result)
	fmt.Println("Hello world" + name)
	const (
		a = iota
		b
		c
	)
	const (
		d = iota
		_
		f
		g
	)
	const (
		FlagA = 1 << iota
		FlagB
		FlagC
		FlagD
	)
	fmt.Println(a, b, c)
	fmt.Println(d, f, g)
	fmt.Println(FlagA, FlagB, FlagC, FlagD)

	switch {
	case hour >= 6:
		fmt.Println("its morning or evening")
	default:
		fmt.Println("Time to work")
	}
	//for...range execution
	message := "hello, Go!"
	for index, runeValue := range message {
		fmt.Printf("Index: %d, Rune: %c \n", index, runeValue)
	}
	// This is an execution of a map
	student := map[string]int{"Alice": 20, "Bob": 22, "Charlie": 21}
	for name, age := range student {
		fmt.Printf("Name: %s, Age: %d \n", name, age)
	}
	//a defer function call example
	fmt.Println("Start")
	//defer fmt.Println("Middle")
	fmt.Println("End")
	//with panic
	panicExample()
	fmt.Println("Intruder flagged")
	//better panic example
	//result := divide(7, 0)
	//fmt.Println("Result: ", result)

	example()
	fmt.Println("Continue")
	fmt.Printf("Well here we go again.\n")
	recent := "where were we?"
	fmt.Println(recent)

	//anonymous function
	add := func(m, n int) int {
		return m + n
	}
	resultss := add(5, 3)
	fmt.Println("Result: ", resultss)
	// Calling anon fn directly
	final := func(ab, cd int) int {
		return ab * cd
	}(2, 3) //appending parentheses after the func literal
	fmt.Println("Result: ", final)
	//Closure
	c1 := counter()
	fmt.Println(c1())
	fmt.Println(c1())
	c2 := counter()
	fmt.Println(c2())
	fmt.Println(c1())
	//Factorial using recursion
	fac := factorial(6)
	fmt.Println("Factorial: ", fac)

	//Fibonacci sequence using recursion
	for a3 := range 10 {
		fmt.Println("Fib sequence: ", fib(a3))
	}
}
