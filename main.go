package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

func RutinInt(ch chan int) {
	t := 0
	for {
		t = rand.IntN(10)
		ch <- t
		time.Sleep(300 * time.Millisecond)
	}
}

func RutinString(ch chan string) {
	i := 0
	for {
		ch <- "hi" + strconv.Itoa(i)
		i++

		time.Sleep(1 * time.Second)
	}
}

func RutinFloat(ch chan float64) {
	t := 0.0
	for {
		t = rand.Float64()
		ch <- t
		time.Sleep(5 * time.Second)
	}
}

func main() {
	ch := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan float64)

	go RutinInt(ch)
	go RutinFloat(ch3)
	go RutinString(ch2)

	for {
		select {
		case number := <-ch:
			fmt.Println("intCh", number)
		case str := <-ch2:
			fmt.Println("intCh", str)
		case float := <-ch3:
			fmt.Println("intCh", float)
		}
	}

}
