package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var read = bufio.NewReader(os.Stdin)
	var sum float64
	var grade float64

	for i := 0; i < 20; i++ {
		text, _, _ := read.ReadLine()
		tokens := strings.Split(string(text), " ")

		pointFromGrade, err := getPointFromGrade(tokens[2])

		if err != nil {
			continue
		}

		point := getPointFromString(tokens[1])
		sum += point
		grade += getMultiplyGradeWithPoint(pointFromGrade, point)
	}

	fmt.Println(grade / sum)
}

func getPointFromString(point string) float64 {
	s, _ := strconv.ParseFloat(point, 64)
	return s
}

func getMultiplyGradeWithPoint(grade float64, point float64) float64 {
	return grade * point
}

func getPointFromGrade(grade string) (float64, error) {

	if grade == "P" {
		return 0.0, errors.New("impossible to calculate")
	}

	switch grade {
	case "A+":
		return 4.5, nil
	case "A0":
		return 4.0, nil
	case "B+":
		return 3.5, nil
	case "B0":
		return 3.0, nil
	case "C+":
		return 2.5, nil
	case "C0":
		return 2.0, nil
	case "D+":
		return 1.5, nil
	case "D0":
		return 1.0, nil
	case "F":
		return 0.0, nil
	}

	return 0.0, nil
}
