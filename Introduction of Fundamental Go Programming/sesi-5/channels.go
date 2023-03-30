package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	// mengirim data kepada channel
	// c <- value

	// menerima data dari channel
	// result := <- c

	go introduce("Huda", c)
	go introduce("Satrio", c)
	go introduce("Ardi", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	// channels with anonymous function
	students := []string{"Huda", "Satrio", "Ardi"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)

	// unbuffered channel
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	close(c1)
	time.Sleep(time.Second)

	// buffered channel
	c2 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}

		close(c)
	}(c2)

	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c2 { // v = <- c2
		fmt.Println("main goroutine received value from channel:", v)
	}

	// select
	c3 := make(chan string)
	c4 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c3 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c4 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c3:
			fmt.Println("Received", msg1)
		case msg2 := <-c4:
			fmt.Println("Received", msg2)
		}
	}
}

// channel directions (optinal for type safety)
// c chan string (parameter c is used for sending and receiving data)
// c chan<- string (parameter c is just used for sending data.)
// c <-chan string (parameter c is just used for receiving data)

func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)

	c <- result
}
