package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

const(
	A = iota
	B
	C
)

const typedHello string = "Hello, World"

func main() {

	var test string

	fmt.Scan(&test)

	res, _ := http.Get(test)

	page, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Printf("%s",page)
}