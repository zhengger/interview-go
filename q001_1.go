package main

import (
	"fmt"
	"sync"
)

func Print12AB(number, letter chan bool, wait *sync.WaitGroup) {
	// letter, number := make(chan bool), make(chan bool)
	// wait.Add(1)
	go printNumber(number, letter, wait)
	go printLetter(number, letter, wait)
	number <- true

}

var printNumber = func(number, letter chan bool, wait *sync.WaitGroup) {
	// defer wait.Done()
	i := 1
	for range number {
		fmt.Print(i)
		i++
		fmt.Print(i)
		i++
		letter <- true
	}

}

var printLetter = func(number, letter chan bool, wait *sync.WaitGroup) {
	l := 'A'
	// defer wait.Done()
	for range letter {
		fmt.Print(string(l))
		l++
		fmt.Print(string(l))
		l++
		if l >= 'Z' {
			// fmt.Print("Done")
			wait.Done()
		}
		number <- true
	}

}

func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	wait.Add(1)
	Print12AB(letter, number, &wait)
	wait.Wait()
}
