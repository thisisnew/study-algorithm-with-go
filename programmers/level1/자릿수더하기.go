package main

func main() {

}

func 자릿수더하기(n int) int {

	answer := 0

	for {
		if n == 0 {
			break
		}

		answer += n % 10
		n = n / 10
	}

	return answer
}
