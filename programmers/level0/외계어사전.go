package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(외계어사전([]string{"s", "o", "m", "d"}, []string{"moos", "dzx", "smm", "sunmmo", "som"}))
}

func 외계어사전(spell []string, dic []string) int {

	for _, d := range dic {

		if len([]rune(d)) < len(spell) {
			continue
		}

		for i, s := range spell {

			isMatched := strings.Contains(d, s)

			if !isMatched {
				break
			}

			if i < len(spell)-1 {
				continue
			}

			return 1
		}
	}

	return 2
}
