package main

import "time"

func main() {
	c1, c2 := make(chan string), make(chan string)
	defer func() { close(c1); close(c2) }()
	go func(c chan<- string) {
		<-time.After(time.Second * 3)
		c <- "foo"
	}(c1)

	go func(c chan<- string) {
		<-time.After(time.Second)
		c <- "bar"
	}(c2)

	for i := 1; ; i++ {
		select {
		case val := <-c1:
			println("channel 1", val)

		case val := <-c2:
			println("channel 2", val)

		}
		break
	}

}
