package main

import "fmt"

func main() {

	result,even := half(2)

	fmt.Printf("Divided by 2 is %d and is the number is even? %b" ,result,even)

}

func half(intVar int) (int,bool) {

	var result int = intVar / 2

	return result, intVar % 2 == 0

}
