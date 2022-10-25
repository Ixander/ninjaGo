package main

import "fmt"

//var msg string

func init() {
	//msg = "from init()"
}

func main() {

	//fmt.Println(msg)
	//fmt.Println("I am developer")
	//fmt.Println(findMin(6, 5, -2, 10))
	func() {
		fmt.Println("анонімна функція")
	}()

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
}

func increment() func() int {
	//fmt.Println("test")
	count := 0
	return func() int {
		count++
		//fmt.Println("inside anon func")
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}
func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
