package main

import "fmt"

func main() {

	var sum int = 0

	for i := 0; i < 1000; i++ {
		if divisibleByThree(i) || divisibleByFive(i){
			sum += i
		}
	}

	fmt.Printf("The sum is %d", sum)

}

func divisibleByFive(i int) bool {
	return i % 5 == 0
}

func divisibleByThree(i int) bool {
	return i % 3 == 0
}