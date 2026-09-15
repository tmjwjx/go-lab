package main

import (
	"fmt"
	"sync"
)

var numCh = make(chan int)
var letterCh = make(chan rune)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		letterCh <- 1
		<-numCh
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		<-letterCh
		fmt.Printf("%c\n", rune('A'+i-1))
		numCh <- 1
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()
	close(numCh)
	close(letterCh)

}