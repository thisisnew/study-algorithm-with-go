package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	var read = bufio.NewReader(os.Stdin)

	fmt.Fscanln(read, &input)

	nums := strings.Split(input, ":")

	n, _ := strconv.Atoi(nums[0])
	m, _ := strconv.Atoi(nums[1])

	var isSwitched bool

	if n < m {
		isSwitched = true
		temp := n
		n = m
		m = temp
	}

	for i := m; i >= 1; i-- {
		if n%i == 0 && m%i == 0 {
			n = n / i
			m = m / i
			continue
		}

		if i == 1 {
			break
		}
	}

	if isSwitched {
		fmt.Println(fmt.Sprintf("%v:%v", m, n))
	} else {
		fmt.Println(fmt.Sprintf("%v:%v", n, m))
	}
}
