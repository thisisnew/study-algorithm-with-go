package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(최댓값과최솟값("-1 -1"))
}

func 최댓값과최솟값(s string) string {

	sSlice := strings.Split(s, " ")

	sort.Slice(sSlice, func(i, j int) bool {

		ni, _ := strconv.Atoi(sSlice[i])
		nj, _ := strconv.Atoi(sSlice[j])

		return ni < nj

	})

	return fmt.Sprintf("%s %s", sSlice[0], sSlice[len(sSlice)-1])
}
