package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(큰수만들기("1924", 2))
}

func 큰수만들기(number string, k int) string {

	/*class Solution {
		public String solution(String number, int k) {
		StringBuilder sb = new StringBuilder();
		int index = 0;
		int max = 0;
		for(int i=0; i<number.length() - k; i++) {
		max = 0;
		for(int j = index; j<= k+i; j++) {
		if(max < number.charAt(j)-'0') {
		max = number.charAt(j)-'0';
		index = j+1;
	}
	}
		sb.append(max);
	}
		return sb.toString();
	}
	}*/

	var result strings.Builder
	var idx = 0

	for i := 0; i < len([]rune(number)); i++ {
		var max = 0
		for j := idx; j <= k+1; j++ {
			num, _ := strconv.Atoi(number[j : j+1])
			if max < num {
				max = num
				idx = j + 1
			}
		}

		result.WriteString(strconv.Itoa(max))
	}

	return result.String()
}
