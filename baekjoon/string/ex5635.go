package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	var num int
	var read = bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &num)

	var infoMap map[string]string
	var birthArray []string

	for i := 0; i < num; i++ {
		input, _ := read.ReadString('\n')
		infos := strings.Split(input, " ")
		name := infos[0]
		birth := fmt.Sprintf("%s%s%s", infos[3], infos[2], infos[1])

		birthArray = append(birthArray, birth)
		infoMap[birth] = name
	}

	sort.Slice(birthArray, func(i, j int) bool {
		return birthArray[i] < birthArray[j]
	})

	fmt.Println(infoMap[birthArray[0]])
	fmt.Println(infoMap[birthArray[num-1]])
}
