package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	Print12AB()
// }
func main() {
	ch := make(chan int, 1000)
	wg := sync.WaitGroup{}
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	wg.Wait()
	close(ch)
	fmt.Println("ch is closed: ok")
	time.Sleep(time.Second * 100)
}