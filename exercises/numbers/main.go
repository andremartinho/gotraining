package main

import "fmt"

func main() {

	var smallNumber int
	var largeNumber int

	fmt.Println("Give me a small number")
	fmt.Scan(&smallNumber)

	fmt.Println("Give me a large number")
	fmt.Scan(&largeNumber)

	reminder := largeNumber % smallNumber

	fmt.Printf("The reminder is %d" ,reminder)

}