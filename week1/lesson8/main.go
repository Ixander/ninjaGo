package main

import "fmt"

var a int

func main() {
	num := 3
	sunOfNumbers(&num)
	fmt.Print(num)
}

func sunOfNumbers(num *int) {
	*num += 5
	fmt.Print(*num)
}
