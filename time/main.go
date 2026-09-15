package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now =", now)

	t1 := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)

	fmt.Println("t1 =", t1)
	fmt.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())
	fmt.Println(t1.Format("2006-01-02 15:04:05"))
	diff := t2.Sub(t1)
	fmt.Println("diff =", diff)
	fmt.Println(diff.Minutes(), diff.Seconds())

	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t1)
	fmt.Println("now.Unix()", now.Unix())
}
