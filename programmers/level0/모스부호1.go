package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(모스부호1(".... . .-.. .-.. ---"))
}

func 모스부호1(letter string) string {

	morse := map[string]string{
		".-":   "a",
		"-...": "b",
		"-.-.": "c",
		"-..":  "d",
		".":    "e",
		"..-.": "f",
		"--.":  "g",
		"....": "h",
		"..":   "i",
		".---": "j",
		"-.-":  "k",
		".-..": "l",
		"--":   "m",
		"-.":   "n",
		"---":  "o",
		".--.": "p",
		"--.-": "q",
		".-.":  "r",
		"...":  "s",
		"-":    "t",
		"..-":  "u",
		"...-": "v",
		".--":  "w",
		"-..-": "x",
		"-.--": "y",
		"--..": "z",
	}

	var result strings.Builder

	for _, lt := range strings.Split(letter, " ") {
		result.WriteString(morse[lt])
	}

	return result.String()
}
