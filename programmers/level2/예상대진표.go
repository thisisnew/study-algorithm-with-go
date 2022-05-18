package main

import "fmt"

func main() {
	fmt.Println(예상대진표(8, 4, 7))
}

type Tournament struct {
	PlayerOne int
	PlayerTwo int
}

func 예상대진표(n int, a int, b int) int {

	var result = 1

	tournaments := generateTournaments(n)
	var nextTournaments []Tournament

	if a > b {
		a, b = b, a
	}

out:
	for {

		t := Tournament{}

		for _, tn := range tournaments {

			if t.PlayerOne > 0 && t.PlayerTwo > 0 {
				nextTournaments = append(nextTournaments, t)
				t = Tournament{}
			}

			if tn.PlayerOne == a && tn.PlayerTwo == b {
				break out
			}

			if tn.PlayerOne == a || tn.PlayerTwo == a {
				a = tn.PlayerTwo / 2

				if a%2 != 0 {
					t.PlayerOne = a
				} else {
					t.PlayerTwo = a
				}

				continue
			}

			if tn.PlayerOne == b || tn.PlayerTwo == b {
				b = tn.PlayerTwo / 2

				if b%2 != 0 {
					t.PlayerOne = b
				} else {
					t.PlayerTwo = b
				}

				continue
			}

			n := tn.PlayerTwo / 2

			if n%2 != 0 && t.PlayerOne == 0 {
				t.PlayerOne = n
				continue
			}

			if n%2 == 0 && t.PlayerTwo == 0 {
				t.PlayerTwo = n
				continue
			}
		}

		if t.PlayerOne > 0 && t.PlayerTwo > 0 {
			nextTournaments = append(nextTournaments, t)
			t = Tournament{}
		}

		result++
		tournaments = nextTournaments
		nextTournaments = []Tournament{}
	}

	return result
}

func generateTournaments(n int) []Tournament {

	var result []Tournament

	t := Tournament{}

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			t.PlayerOne = i
		} else {
			t.PlayerTwo = i
			result = append(result, t)
			t = Tournament{}
		}
	}

	return result
}
