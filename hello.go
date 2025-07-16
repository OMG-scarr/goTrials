/*
	This is my 1st go project

File: hello.go
Description: This program prints "Hello World" to the console and that's just the start
Author: Gatheru Oscar
Date: 2025-07-13
*/

package main

import (
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
func main() {
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
}
