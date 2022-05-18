package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(조이스틱("JEROEN"))
}

func 조이스틱(name string) int {
	var result int
	var idx int
	var prevIdx []int
	var isLast bool

	for {

		if isLast {
			break
		}

		token := name[idx : idx+1]
		tokenToInt := int(big.NewInt(0).SetBytes([]byte(token)).Uint64()) - 65

		upCnt := controlUp(tokenToInt)
		downCnt := controlDown(tokenToInt)

		if upCnt <= downCnt {
			result += upCnt
		} else {
			result += downCnt
		}

		prevIdx = append(prevIdx, idx)
		lCnt, lIdx := cursorLeft(idx, prevIdx, name)
		rCnt, rIdx := cursorRight(idx, prevIdx, name)

		isLast = isLastIdx(idx, len(name)-1)

		if rCnt <= lCnt {
			result += rCnt
			idx = rIdx
		} else {
			result += lCnt
			idx = lIdx
		}
	}

	return result
}

func isLastIdx(idx, ln int) bool {
	if idx == ln {
		return true
	}

	return false
}

func controlUp(dist int) int {

	var result int
	var i int

	for {
		if i == dist {
			break
		}

		result++
		i++
	}

	return result
}

func controlDown(dist int) int {

	var result int
	var i int

	for {
		if i == dist {
			break
		}

		result++
		i--

		if i < 0 {
			i = 25
		}
	}

	return result
}

func cursorLeft(idx int, prevIdx []int, name string) (int, int) {

	var result int
	var i = idx

	for {

		result++
		i--

		if i < 0 {
			i = len(name) - 1
		}

		if i == idx {
			break
		}

		if name[i:i+1] == "A" {
			continue
		}

		if isInPrevIdx(i, prevIdx) {
			continue
		}

		break
	}

	return result, i
}

func cursorRight(idx int, prevIdx []int, name string) (int, int) {

	var result int

	for i := idx + 1; i < len(name); i++ {

		result++
		idx++

		if name[i:i+1] == "A" {
			continue
		}

		if isInPrevIdx(idx, prevIdx) {
			continue
		}

		return result, idx
	}

	return result, idx
}

func isInPrevIdx(idx int, prevIdx []int) bool {

	for _, i := range prevIdx {
		if idx == i {
			return true
		}
	}

	return false
}
