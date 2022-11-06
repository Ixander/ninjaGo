package main

import (
	"fmt"
	"reflect"
)

func main() {
	var q = new(int)
	e := 10
	fmt.Println(reflect.TypeOf(q))
	fmt.Println(reflect.TypeOf(e))

	fmt.Println(&e)
}
