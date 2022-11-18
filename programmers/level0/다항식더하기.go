package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(다항식더하기("x + x + x"))
}

func 다항식더하기(polynomial string) string {

	var xNum int
	var cons int
	pSl := strings.Split(polynomial, "+")

	for _, p := range pSl {
		p = strings.TrimSpace(p)

		if strings.Contains(p, "x") {
			v := p
			v = strings.ReplaceAll(v, "x", "")

			var n int
			if v == "" {
				n = 1
			} else {
				n, _ = strconv.Atoi(v)
			}

			xNum += n
			continue
		}

		n, _ := strconv.Atoi(p)
		cons += n
	}

	if xNum > 0 && cons > 0 {
		if xNum == 1 {
			return fmt.Sprintf("x + %v", cons)
		}

		return fmt.Sprintf("%vx + %v", xNum, cons)
	}

	if xNum > 0 {
		if xNum == 1 {
			return fmt.Sprintf("x")
		}

		return fmt.Sprintf("%vx", xNum)
	}

	if cons > 0 {
		return fmt.Sprintf("%v", cons)
	}

	return ""
}
