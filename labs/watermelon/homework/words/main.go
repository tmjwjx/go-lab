package main

import "fmt"

func filter(words []string, handler func(string) bool) (ans []string) {
	for _, word := range words {
		if handler(word) {
			ans = append(ans, word)
		}
	}
	return
}

func judgmentEven(word string) bool {
	if len(word)%2 == 0 {
		return true
	}
	return false
}

func containA(word string) bool {
	for _, val := range word {
		if val == 'a' {
			return true
		}
	}
	return false
}

func main() {
	words := []string{
		"aaaa",
		"bbb",
		"a",
		"bbba",
	}
	for _, word := range filter(words, judgmentEven) {
		fmt.Println(word)
	}
	fmt.Println()
	for _, word := range filter(words, containA) {
		fmt.Println(word)
	}
}