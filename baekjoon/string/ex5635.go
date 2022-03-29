package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var infoMap = map[string]string{}
	var birthArray []string

	reader := bufio.NewReader(os.Stdin)
	var num int

	fmt.Fscan(reader, &num)

	for i := 0; i < num; i++ {
		text, _, _ := reader.ReadLine()

		fmt.Println(text)

		input := string(text)
		infos := strings.Split(input, " ")
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
