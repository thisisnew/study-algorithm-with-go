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

		if input == "END" {
			break
		}

		var vowelCnt int
		var consCnt int
		var prevRn = input[0:1]

		for i := 1; i < len(input); i++ {
			ch := input[i : i+1]

			switch {
			case ch == "a", ch == "i", ch == "o", ch == "u", ch == "e":
				vowelCnt++
			default:
				consCnt++
			}

			if vowelCnt == 3 || consCnt == 3 {
				fmt.Println(fmt.Sprintf("<%s> is not acceptable.", input))
				break
			}

			if prevRn == ch {
				str := fmt.Sprintf("%s%s", prevRn, ch)

				if str != "ee" && str != "oo" {
					fmt.Println(fmt.Sprintf("<%s> is not acceptable.", input))
					break
				}
			}

			prevRn = ch
		}

		if vowelCnt == 0 {
			fmt.Println(fmt.Sprintf("<%s> is not acceptable.", input))
			continue
		}

		fmt.Println(fmt.Sprintf("<%s> is acceptable.", input))
	}
}
