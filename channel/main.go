package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	close(ch)
	v, ok := <-ch
	for ok {
		v, ok = <-ch
	}

	for v, ok = <-ch; ok; {
		fmt.Println(v)
	}

	fmt.Println(v, ok)
}
