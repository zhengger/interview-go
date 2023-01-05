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
	wait.Wait()
}

var printNumber = func(number, letter chan bool, wait *sync.WaitGroup) {
	i := 1
OuterLoop:
	for {
		select {
		case <-number:
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- true
		default:
			break OuterLoop
		}
	}
}

var printLetter = func(number, letter chan bool, wait *sync.WaitGroup) {
	wait.Add(1)
	l := 'A'
OuterLoop:
	for {
		select {
		case <-letter:
			if l >= 'Z' {
				fmt.Print("Done")
				wait.Done()
				break
			}
			fmt.Print(string(l))
			l++
			fmt.Print(string(l))
			l++
			number <- true
		default:
			break OuterLoop
		}
	}
}
