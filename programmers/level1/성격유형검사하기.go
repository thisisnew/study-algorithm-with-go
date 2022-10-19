package main

import (
	"fmt"
	"strings"
)

func main() {

	survey := []string{
		"TR", "RT", "TR",
	}

	choices := []int{
		7, 1, 3,
	}

	fmt.Println(성격유형검사하기(survey, choices))
}

func 성격유형검사하기(survey []string, choices []int) string {

	var result strings.Builder
	var points = map[string]int{
		"R": 0,
		"T": 0,
		"C": 0,
		"F": 0,
		"J": 0,
		"M": 0,
		"A": 0,
		"N": 0,
	}

	for i, s := range survey {

		f := s[:1]
		b := s[1:]

		switch {
		case choices[i] == 1:
			points[f] += 3
		case choices[i] == 2:
			points[f] += 2
		case choices[i] == 3:
			points[f] += 1
		case choices[i] == 4:
			continue
		case choices[i] == 5:
			points[b] += 1
		case choices[i] == 6:
			points[b] += 2
		case choices[i] == 7:
			points[b] += 3
		}

	}

	for _, t := range []string{"RT", "CF", "JM", "AN"} {
		f := t[:1]
		b := t[1:]

		var pointF = points[f]
		var pointB = points[b]

		if pointF > pointB {
			result.WriteString(f)
			continue
		}

		if pointB > pointF {
			result.WriteString(b)
			continue
		}

		if f > b {
			result.WriteString(b)
			continue
		}

		if b > f {
			result.WriteString(f)
			continue
		}
	}

	return result.String()
}
