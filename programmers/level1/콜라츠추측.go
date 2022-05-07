package main

const maximumCount = 500

func main() {

}

func 콜라츠추측(num int) int {

	var cnt int
	var result int

	for {

		if cnt > maximumCount {
			result = -1
			break
		}

		if isEven(num) {
			num = num / 2
		} else {
			num = num*3 + 1
		}

		if num == 1 {
			result = cnt
			break
		}

	}

	return result
}

func isEven(num int) bool {

	if num%2 == 0 {
		return true
	}

	return false
}
