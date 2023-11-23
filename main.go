package main

import (
	"fmt"
	"time"
)

func process() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// T1
func main() {
	// go process() // T2
	// go process() // T3
	// process() // T1

	channel := make(chan int)
	go func() {
		// channel <- 1 // T2
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println("Jogou no canal", i)
		}
	}()
	go func() {
		// channel <- 1 // T2
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println("Jogou no canal", i)
		}
	}()

	// fmt.Println(<-channel) // esvaziar o canal
	// time.Sleep(time.Second * 2)

	// for x := range channel {
	// 	fmt.Println("Recebeu do canal", x)
	// 	time.Sleep(time.Second)
	// }

	go worker(channel, 1)
	worker(channel, 2)
}

func worker(channel chan int, workerID int) {
	for {
		fmt.Println("Recebeu do canal", <-channel, "no worker", workerID)
		time.Sleep(time.Second)
	}
}