package main

import (
	"fmt"
	"time"
)

func main() {

	loc, err := time.LoadLocation("Asia/Seoul")

	if err != nil {
		panic(err)
	}

	t := time.Now().In(loc)
	fmt.Println(t.Year())
	fmt.Println(int(t.Month()))
	fmt.Println(t.Day())
}
