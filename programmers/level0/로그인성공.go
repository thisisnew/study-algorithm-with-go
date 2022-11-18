package main

import "fmt"

func main() {
	fmt.Println(로그인성공([]string{"meosseugi", "1234"}, [][]string{{"rardss", "123"}, {"yyoom", "1234"}, {"meosseugi", "1234"}}))
}

func 로그인성공(id_pw []string, db [][]string) string {

	for _, account := range db {
		if account[0] == id_pw[0] {
			if account[1] == id_pw[1] {
				return "login"
			}

			return "wrong pw"
		}
	}

	return "fail"
}
