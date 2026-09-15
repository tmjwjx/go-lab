package main

import (
	"fmt"
)

func SliceRise(s *[]int) {
	*s = append(*s, 0)
	for i := range *s {
		(*s)[i]++
	}
}

func SlicePrint() {
	s1 := []int{1, 2}
	fmt.Println(s1)
	fmt.Println(&s1)
	s2 := s1
	fmt.Println(s1, s2)
	fmt.Println(&s1, &s2)
	s2 = append(s2, 3)
	fmt.Println(s1, s2)
	fmt.Println(&s1, &s2)
	SliceRise(&s1)
	SliceRise(&s2)
	fmt.Println(s2[3])
	fmt.Println(s1, s2)
	fmt.Println(&s1, &s2)

}

func main() {
	SlicePrint()
}
