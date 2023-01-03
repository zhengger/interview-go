package main

import (
	"fmt"
	"sync"
)

func Print12AB() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go printNumber(number, letter, &wait)
	go printLetter(number, letter, &wait)
	number <- true
}

var printNumber = func(number, letter chan bool, wait *sync.WaitGroup) {
	// wait.Add(1)
	//
	i := 1
	for {
		select {
		case <-number:
			fmt.Print(i)
			i++
			fmt.Print(i)
			letter <- true
		}
	}
}

var printLetter = func(number, letter chan bool, wait *sync.WaitGroup) {
	wait.Add(1)
	l := 'A'
	for {
		select {
		case <-letter:
			if l >= 'Z' {
				fmt.Print("Done")
				wait.Done()
			}
			fmt.Print(string(l))
			l++
			fmt.Print(string(l))
			number <- true
		}
	}
}

func main() {
	Print12AB()
	// time.Sleep(3 * time.Second)
}
