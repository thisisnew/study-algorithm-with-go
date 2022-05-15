package main

import (
	"fmt"
	"sort"
)

type Students struct {
	num     int
	lost    bool
	reserve bool
}

func main() {

	n := 3
	lost := []int{1, 2}
	reserve := []int{2, 3}

	fmt.Println(체육복(n, lost, reserve))
}

func 체육복(n int, lost []int, reserve []int) int {

	sort.Slice(lost, func(i, j int) bool {
		return lost[i] < lost[j]
	})

	sort.Slice(reserve, func(i, j int) bool {
		return reserve[i] < reserve[j]
	})

	var lostAndNotReserved []int
	var lostAndNotBorrowed []int

	for i := 0; i < len(lost); i++ {

		isReserved, reservedButNotLost := isReservedStudent(reserve, lost[i])

		reserve = reservedButNotLost

		if isReserved {
			continue
		}

		lostAndNotReserved = append(lostAndNotReserved, lost[i])
	}

	for i := 0; i < len(lostAndNotReserved); i++ {
		canBorrow, notReservedYet := canBorrowGymSuit(reserve, lostAndNotReserved[i])

		reserve = notReservedYet

		if canBorrow {
			continue
		}

		lostAndNotBorrowed = append(lostAndNotBorrowed, lostAndNotReserved[i])
	}

	return n - len(lostAndNotBorrowed)
}

func isReservedStudent(reserve []int, n int) (bool, []int) {

	var reservedButNotLost []int
	var isReserved bool

	for _, rn := range reserve {
		if rn == n {
			isReserved = true
			continue
		}

		reservedButNotLost = append(reservedButNotLost, rn)
	}

	return isReserved, reservedButNotLost
}

func canBorrowGymSuit(reserve []int, n int) (bool, []int) {

	var notLendYet []int
	var canBorrow bool

	for _, rn := range reserve {
		if rn == n-1 || rn == n+1 {
			canBorrow = true
			continue
		}

		notLendYet = append(notLendYet, rn)
	}

	return canBorrow, notLendYet
}
