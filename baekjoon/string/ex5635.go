package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var infoMap = map[string]string{}
	var birthArray []string

	var num int
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		input := string(line)
		infos := strings.Split(input, " ")

		if len(infos) < 3 {
			n, _ := strconv.Atoi(input)
			num = n
			continue
		}

		name := infos[0]
		birth := fmt.Sprintf("%s%s%s", infos[3], addZero(infos[2]), addZero(infos[1]))

		birthArray = append(birthArray, birth)
		infoMap[birth] = name
	}

	sort.Slice(birthArray, func(i, j int) bool {
		return birthArray[i] < birthArray[j]
	})

	fmt.Println(infoMap[birthArray[0]])
	fmt.Println(infoMap[birthArray[num-1]])
}

func addZero(num string) string {

	n, _ := strconv.Atoi(num)

	if n < 10 {
		num = "0" + num
	}

	return num
}
