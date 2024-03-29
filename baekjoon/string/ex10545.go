package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var usualKeyMaps = getUsuallyKeymaps()

func main() {
	var read = bufio.NewReader(os.Stdin)
	text, _, _ := read.ReadLine()
	nums := strings.Split(string(text), " ")
	keymaps := brokenKeymaps(nums)

	var input string
	fmt.Fscanln(read, &input)

	fmt.Println(getKeyMapTouches(input, keymaps))
}

func getUsuallyKeymaps() map[string][]string {

	result := map[string][]string{}

	for i := 1; i <= 9; i++ {
		var props []string

		switch i {
		case 1:
		case 2:
			props = []string{"a", "b", "c"}
		case 3:
			props = []string{"d", "e", "f"}
		case 4:
			props = []string{"g", "h", "i"}
		case 5:
			props = []string{"j", "k", "l"}
		case 6:
			props = []string{"m", "n", "o"}
		case 7:
			props = []string{"p", "q", "r", "s"}
		case 8:
			props = []string{"t", "u", "v"}
		case 9:
			props = []string{"w", "x", "y", "z"}
		}

		result[strconv.Itoa(i)] = props
	}

	return result
}

func brokenKeymaps(nums []string) map[string][]string {

	result := map[string][]string{}

	for i, num := range nums {
		result[strconv.Itoa(i+1)] = usualKeyMaps[num]
	}

	return result
}

func getKeyMapTouches(input string, keymaps map[string][]string) string {

	var result strings.Builder
	var prevWord string

	for i := 0; i < len(input); i++ {
		var s = input[i : i+1]

		for key, val := range keymaps {

			if !isTokenInKeymaps(s, val) {
				continue
			}

			for i, v := range val {
				if v != s {
					continue
				}

				if prevWord == key {
					result.WriteString("#")
				}

				for j := 0; j < i+1; j++ {
					result.WriteString(key)
				}

				prevWord = key
			}
		}
	}

	return result.String()
}

func isTokenInKeymaps(s string, val []string) bool {
	if len(val) == 0 || strings.Compare(s, val[len(val)-1]) == 1 {
		return false
	}

	return true
}
