package main

import "fmt"

func main() {

	survey := []string{
		"",
	}

	choices := []int{
		-1, 1, -1, 1,
	}

	fmt.Println(성격유형검사하기(survey, choices))
}

func 성격유형검사하기(survey []string, choices []int) string {

	var result string
	var points = map[string]int{}

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

	return result
}
