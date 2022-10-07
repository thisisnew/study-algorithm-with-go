package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var read = bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(read, &n, &m)

	var groupNames = map[string][]string{}
	var members = map[string]string{}

	for i := 0; i < n; i++ {
		var groupName string
		fmt.Fscanln(read, &groupName)

		var memberCount int
		fmt.Fscanln(read, &memberCount)

		var memberSl = make([]string, memberCount)
		for j := 0; j < memberCount; j++ {
			var member string
			fmt.Fscanln(read, &member)
			members[member] = groupName
			memberSl[j] = member
		}

		groupNames[groupName] = memberSl
	}

	for i := 0; i < m; i++ {
		var name string
		fmt.Fscanln(read, &name)

		var qType int
		fmt.Fscanln(read, &qType)

		switch qType {
		case 0:
			printGroupMembers(groupNames[name])
		case 1:
			printGroupName(members[name])
		}
	}

}

func printGroupMembers(members []string) {

	sort.Slice(members, func(i, j int) bool {
		return members[i] < members[j]
	})

	for _, member := range members {
		fmt.Println(member)
	}
}

func printGroupName(groupName string) {
	fmt.Println(groupName)
}
