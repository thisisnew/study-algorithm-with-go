package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var n int
	var m int
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &n, &m)

	noHearPersons := make(map[string]bool, n)
	var noHearNoLookPerson []string
	var name string

	for i := 0; i < n; i++ {
		fmt.Fscanln(read, &name)

		noHearPersons[name] = true
	}

	for i := 0; i < m; i++ {
		fmt.Fscanln(read, &name)

		if _, ok := noHearPersons[name]; !ok {
			continue
		}

		noHearNoLookPerson = append(noHearNoLookPerson, name)
	}

	fmt.Println(len(noHearNoLookPerson))

	sort.Slice(noHearNoLookPerson, func(i, j int) bool {
		return noHearNoLookPerson[i] < noHearNoLookPerson[j]
	})

	for _, name := range noHearNoLookPerson {
		fmt.Println(name)
	}
}
