package main

import "fmt"

func main() {
	fmt.Println(교점에별만들기([][]int{{2, -1, 4}, {-2, -1, 4}, {0, -1, 1}, {5, -8, -12}, {5, 8, 12}}))
}

func 교점에별만들기(line [][]int) []string {

	var meets [][]float64

	for i := 0; i < len(line)-1; i++ {
		a1 := line[i][0]
		b1 := line[i][1]
		c1 := line[i][2]

		for j := i + 1; j < len(line); j++ {
			a2 := line[j][0]
			b2 := line[j][1]
			c2 := line[j][2]

			x := float64(b1*c2-b2*c1) / float64(a1*b2-a2*b1)
			y := float64(-1)*float64(a1/b1)*x - float64(c1/b1)

			meets = append(meets, []float64{x, y})
		}
	}

	fmt.Println(meets)

	return []string{}
}
