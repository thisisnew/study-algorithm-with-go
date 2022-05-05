package main

import (
	"fmt"
	"time"
)

const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(year2016(5, 24))
}

func year2016(a int, b int) string {

	dateString := fmt.Sprintf("2016-%v-%v", a, b)

	date, _ := time.Parse("2006-01-02", dateString)

	return getDayOfWeek(date.Day())

}

func getDayOfWeek(day int) string {

	switch day {
	case Sunday:
		return "SUN"
	case Monday:
		return "MON"
	case Tuesday:
		return "TUE"
	case Wednesday:
		return "WED"
	case Thursday:
		return "THU"
	case Friday:
		return "FRI"
	case Saturday:
		return "SAT"
	}

	return ""
}
