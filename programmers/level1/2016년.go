package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(year2016(5, 24))
}

func year2016(a int, b int) string {

	dateString := fmt.Sprintf("2016-%v-%v", addZero(a), addZero(b))

	date, _ := time.Parse("2006-01-02", dateString)

	return strings.ToUpper(date.Weekday().String()[:3])

}

func addZero(n int) string {

	if n < 10 {
		return fmt.Sprintf("0%v", n)
	}

	return fmt.Sprintf("%v", n)
}
