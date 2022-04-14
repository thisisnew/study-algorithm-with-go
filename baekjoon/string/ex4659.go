package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var read = bufio.NewReader(os.Stdin)

	for {
		var input string
		fmt.Fscanln(read, &input)

		if input == "end" {
			break
		}

		var vowelCnt int
		var consCnt int

		var prevRn = input[0:1]
		var isAcceptable = true

		for i := 0; i < len(input); i++ {
			ch := input[i : i+1]

			switch {
			case ch == "a", ch == "i", ch == "o", ch == "u", ch == "e":
				vowelCnt++

			default:
				consCnt++
			}

			if vowelCnt == 3 || consCnt == 3 {
				isAcceptable = false
				break
			}

			if i > 0 {

				if prevRn == ch {
					str := fmt.Sprintf("%s%s", prevRn, ch)

					if str != "ee" && str != "oo" {
						isAcceptable = false
						break
					}
				}

				prevRn = ch
			}
		}

		if isAcceptable && vowelCnt == 0 {
			isAcceptable = false
		}

		if !isAcceptable {
			fmt.Println(fmt.Sprintf("<%s> is not acceptable.", input))
		} else {
			fmt.Println(fmt.Sprintf("<%s> is acceptable.", input))
		}

	}
}
