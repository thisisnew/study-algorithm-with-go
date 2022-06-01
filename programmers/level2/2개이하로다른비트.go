package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numbers := []int64{2, 7}

	fmt.Println(두개이하로다른비트(numbers))
}

func 두개이하로다른비트(numbers []int64) []int64 {

	var result = make([]int64, len(numbers))

	for i, n := range numbers {
		result[i] = getMinimalNumberHasDifBits(n)
	}

	return result
}

func getMinimalNumberHasDifBits(n int64) int64 {

	var result = n
	binN := strconv.FormatInt(n, 2)

	for {

		result++
		binDif := strconv.FormatInt(result, 2)

		if countDifBits(setBitsLength(binN, binDif)) <= 2 {
			break
		}

	}

	return result
}

func setBitsLength(binN, binDif string) (string, string) {

	if len(binN) > len(binDif) {
		return binN, addZeros(binDif, len(binN)-len(binDif))
	}

	if len(binN) < len(binDif) {
		return addZeros(binN, len(binDif)-len(binN)), binDif
	}

	return binN, binDif
}

func addZeros(bin string, n int) string {

	var result strings.Builder

	for i := 0; i < n; i++ {
		result.WriteString("0")
	}

	result.WriteString(bin)

	return result.String()
}

func countDifBits(binN, binDif string) int {

	var result int

	for i := 0; i < len(binDif); i++ {

		if binDif[i:i+1] == binN[i:i+1] {
			continue
		}

		result++
	}

	return result
}
