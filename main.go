package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) {

		for {
			ms := (rand.IntN(10) + 1) * 100
			i := rand.IntN(10)
			time.Sleep(time.Duration(ms) * time.Millisecond)
			fmt.Println("1 рутина Поспал", ms, "Миллисекунд")
			ch <- i
		}
	}(ch1)

	go func(ch chan int) {

		for {
			ms := (rand.IntN(10) + 1) * 100
			i := rand.IntN(10)
			time.Sleep(time.Duration(ms) * time.Millisecond)
			fmt.Println("2 рутина Поспал", ms, "Миллисекунд")
			ch <- i
		}
	}(ch2)

	time.Sleep(500 * time.Millisecond)

	for {
		select {
		case v1 := <-ch1:
			fmt.Println("Сработала секция 1", v1)
		case v2 := <-ch2:
			fmt.Println("Сработала секция 2", v2)
		}
	}
}
