package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(개인정보수집유효기간("2022.05.19", []string{"A 6", "B 12", "C 3"}, []string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"}))
}

func 개인정보수집유효기간(today string, terms []string, privacies []string) []int {

	termsMap := getValidityMap(terms)
	var result []int

	for i, privacy := range privacies {
		collectionDate, termsType := privacyEntities(privacy)

		if isAbrogateDate(today, geExpirationDate(collectionDate, termsMap[termsType])) {
			result = append(result, i+1)
		}
	}

	return result
}

func getValidityMap(terms []string) map[string]int {
	var result = make(map[string]int)

	for _, term := range terms {
		sp := strings.Split(term, " ")
		month, _ := strconv.Atoi(sp[1])
		result[sp[0]] = month
	}

	return result
}

func privacyEntities(privacy string) (string, string) {
	pr := strings.Split(privacy, " ")
	return pr[0], pr[1]
}

func geExpirationDate(collectionDate string, termsDuration int) string {
	result, _ := time.Parse("2006.01.02", collectionDate)
	return result.AddDate(0, termsDuration, 0).Format("2006.01.02")
}

func getDayEntities(day string) (string, string, string) {
	result := strings.Split(day, ".")
	return result[0], result[1], result[2]
}

func isAbrogateDate(today string, expireDate string) bool {
	expireYear, expireMonth, expireDay := getDayEntities(expireDate)
	tYear, tMonth, tDay := getDayEntities(today)

	if tYear < expireYear {
		return false
	}

	if tYear == expireYear && tMonth < expireMonth {
		return false
	}

	if tYear == expireYear && tMonth == expireMonth && tDay < expireDay {
		return false
	}

	return true
}
