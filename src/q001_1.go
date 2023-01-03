package main

import (
	"fmt"
	"sync"
)

func Print12AB(wait *sync.WaitGroup) {
	letter, number := make(chan bool), make(chan bool)


	go printNumber(number, letter, wait)
	go printLetter(number, letter, wait)
	number <- true
}

var printNumber = func(number, letter chan bool, wait *sync.WaitGroup) {
	wait.Add(1)
	i := 1
	for range number {
		fmt.Print(i)
		i++
		fmt.Print(i)
		letter <- true
	}
	wait.Done()
}

var printLetter = func(number, letter chan bool, wait *sync.WaitGroup) {
	wait.Add(1)
	l := 'A'
	for range letter {
		if l > 'Z' {
			fmt.Print("Done")
			wait.Done()
		}
		fmt.Print(string(l))
		l++
		fmt.Print(string(l))
		number <- true
	}
	wait.Done()
}

func main() {
	wait := sync.WaitGroup{}
	Print12AB(&wait)
	wait.Wait()
}
