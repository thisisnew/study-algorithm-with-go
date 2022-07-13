package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var sl []string
	var endCnt = 0

	for {
		var f string
		fmt.Fscanln(read, &f)

		if f == "END" {
			endCnt++
		} else {
			endCnt = 0
		}

		sl = append(sl, f)

		if endCnt == 2 {
			break
		}

	}

	var i = 0
	var c = 1
	for i < len(sl) {

		var f = sl[i]
		var s = sl[i+1]

		if f == "END" && s == "END" {
			break
		}

		if isWithdrawRight(f, s) {
			fmt.Println(fmt.Sprintf("Case %v: same", c))
		} else {
			fmt.Println(fmt.Sprintf("Case %v: different", c))
		}

		c++
		i += 2
	}

}

func isWithdrawRight(f, s string) bool {

	if len([]rune(f)) != len([]rune(s)) {
		return false
	}

	if sortString(f) != sortString(s) {
		return false
	}

	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
