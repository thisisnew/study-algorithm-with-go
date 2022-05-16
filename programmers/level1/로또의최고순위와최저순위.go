package main

import (
	"fmt"
)

func main() {
	lottos := []int{45, 4, 35, 20, 3, 9}

	win_nums := []int{20, 9, 3, 45, 4, 35}

	fmt.Println(로또의최고순위와최저순위(lottos, win_nums))
}

func 로또의최고순위와최저순위(lottos []int, win_nums []int) []int {

	count := getMatchedNum(&lottos, &win_nums)
	highest := count
	lowest := count

	for _, lt := range lottos {
		if lt == 0 {
			highest++
		}
	}

	return []int{getRank(highest), getRank(lowest)}
}

func getMatchedNum(lottos *[]int, win_nums *[]int) int {
	var result int

	for i := 0; i < len(*lottos); i++ {
		for j := 0; j < len(*win_nums); j++ {
			if (*lottos)[i] == (*win_nums)[j] {
				result++
				(*lottos)[i] = -1
				(*win_nums)[j] = -1
				break
			}
		}
	}

	return result
}

func getRank(n int) int {

	switch n {
	case 6:
		return 1
	case 5:
		return 2
	case 4:
		return 3
	case 3:
		return 4
	case 2:
		return 5
	default:
		return 6
	}

}
