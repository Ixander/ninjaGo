package main

import "fmt"

func main() {
	massage := "Скоро я стану Ніндзя!"
	printMessage(&massage)

	fmt.Println(massage)
}

func printMessage(message *string) {
	*message += " (з функції printMessage())"
	fmt.Println(*message)
}
