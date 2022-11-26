package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(베스트앨범([]string{"classic", "pop", "classic", "classic", "pop"}, []int{500, 600, 150, 800, 2500}))
}

func 베스트앨범(genres []string, plays []int) []int {

	genresPlaysMap, genresIdxMap, totalPlaysMap := getGenresPlaysMap(genres, plays)
	totalPlays := getTotalPlaysDesc(totalPlaysMap)
	var result []int

	for _, play := range totalPlays {
		result = append(result, addBestAlbum(getGenresMatchesPlays(totalPlaysMap, play), genresPlaysMap, genresIdxMap)...)
	}

	return result
}

func getGenresPlaysMap(genres []string, plays []int) (map[string][]int, map[int][]int, map[string]int) {
	var genresPlaysMap = map[string][]int{}
	var genresIdxMap = map[int][]int{}
	var totalPlaysMap = map[string]int{}

	for i, g := range genres {
		genresPlaysMap[g] = append(genresPlaysMap[g], plays[i])
		genresIdxMap[plays[i]] = append(genresIdxMap[plays[i]], i)
		totalPlaysMap[g] += plays[i]
	}

	return genresPlaysMap, genresIdxMap, totalPlaysMap
}

func getTotalPlaysDesc(totalPlaysMap map[string]int) []int {

	var plays []int

	for _, v := range totalPlaysMap {
		plays = append(plays, v)
	}

	sort.Slice(plays, func(i, j int) bool {
		return plays[i] > plays[j]
	})

	return plays
}

func getGenresMatchesPlays(totalPlaysMap map[string]int, play int) []string {
	var result []string

	for k, v := range totalPlaysMap {
		if v == play {
			result = append(result, k)
		}
	}

	return result
}

func addBestAlbum(genresMatchesPlays []string, genresPlaysMap map[string][]int, genresIdxMap map[int][]int) []int {

	var result []int

	if len(genresMatchesPlays) == 1 {
		ps := genresPlaysMap[genresMatchesPlays[0]]

		if len(ps) == 1 {
			result = append(result, genresIdxMap[ps[0]][0])
		} else {
			sort.Slice(ps, func(i, j int) bool {
				return ps[i] > ps[j]
			})

			ps0 := ps[0]
			ps1 := ps[1]

			if ps0 == ps1 {
				result = append(result, genresIdxMap[ps0][0])
				result = append(result, genresIdxMap[ps0][1])
			} else {
				result = append(result, genresIdxMap[ps0][0])
				result = append(result, genresIdxMap[ps1][0])
			}

		}
	}

	return result
}
