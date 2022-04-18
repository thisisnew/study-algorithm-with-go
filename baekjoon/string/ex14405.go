package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscanln(read, &input)

	var yes = true
	var ans string
	for _, rn := range input {
		ans += string(rn)

		if len(ans) == 2 {
			if ans == "pi" || ans == "ka" {
				ans = ""
				continue
			}
		}

		if len(ans) == 3 {
			if ans == "chu" {
				ans = ""
				continue
			}
		}

		if len(ans) > 3 {
			yes = false
			break
		}
	}

	if !yes {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
