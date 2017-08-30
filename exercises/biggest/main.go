package main

import "fmt"

func main() {
	fmt.Printf("The biggest number is %d", biggest(100,2,3,4,20,-1))
}

func biggest(numbers ...int) int {

	var biggestNumber int = 0

	for _, value := range numbers{
		if value > biggestNumber {
			biggestNumber = value
		}
	}

	return biggestNumber
}

