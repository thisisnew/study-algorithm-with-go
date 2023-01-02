package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(binaryGap(32))
}

func binaryGap(N int) int {

	var result = 0
	var binNum = strconv.FormatInt(int64(N), 2)

	if strings.Count(binNum, "1") < 2 {
		return result
	}

	//var idx = strings.Index(binNum, "1")

	//for {
	//
	//}

	return result

}
