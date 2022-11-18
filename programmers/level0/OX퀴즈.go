package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(OX퀴즈([]string{"3 - 4 = -3", "5 + 6 = 11"}))
}

func OX퀴즈(quiz []string) []string {

	var result []string

	for _, q := range quiz {

		qSl := strings.Split(q, " ")
		var sum, _ = strconv.Atoi(qSl[0])
		var isPlus bool

		for i := 1; i < len(qSl); i++ {
			v := qSl[i]

			if v == "+" {
				isPlus = true
				continue
			}

			if v == "-" {
				isPlus = false
				continue
			}

			if v == "=" {
				ans, _ := strconv.Atoi(qSl[len(qSl)-1])

				if sum == ans {
					result = append(result, "O")
				} else {
					result = append(result, "X")
				}

				break
			}

			n, _ := strconv.Atoi(v)

			if isPlus {
				sum += n
			} else {
				sum -= n
			}
		}
	}

	return result
}
