package main

func main() {

}

func 괄호회전하기(s string) int {
	return -1
}

func isValidBracedString(s string) bool {

	firstLetter := s[0:1]
	lastLetter := s[len(s)-1:]

	switch firstLetter {
	case "}", "]", ")":
		return false
	}

	switch lastLetter {
	case "{", "[", "(":
		return false
	}

	return true
}
