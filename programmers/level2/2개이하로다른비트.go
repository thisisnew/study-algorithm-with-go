package main

import "strconv"

func main() {

}

func 두개이하로다른비트(numbers []int64) []int64 {

	var result = make([]int64, len(numbers))

	for i, n := range numbers {
		minimalN := getMinimalNumberHasDifBits(n)

		result[i] = minimalN
	}

	return result
}

func getMinimalNumberHasDifBits(n int64) int64 {

	var result = n
	binN := strconv.FormatInt(n, 2)

	for {

		binDif := strconv.FormatInt(result, 2)

		if countDifBits(binN, binDif) <= 2 {
			break
		}

		result++
	}

	return result

}

func countDifBits(binN, binDif string) int {

	var result int

	for i := 0; i < len(binDif); i++ {

		if binDif[i:i+1] != binN[i:i+1] {
			result++
		}

	}

	return result

}
