package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if divisibleByThree(i) && divisibleByFive(i){
			fmt.Println("FizzBuzz")
		}else if divisibleByFive(i){
			fmt.Println("Buzz")
		}else if divisibleByThree(i){
			fmt.Println("Fizz")
		}
	}
}
func divisibleByFive(i int) bool {
	return i % 5 == 0
}

func divisibleByThree(i int) bool {
	return i % 3 == 0
}
