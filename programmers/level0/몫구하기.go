package level0

import "fmt"

func main() {
	fmt.Println(몫구하기(10, 5))
}

func 몫구하기(num1 int, num2 int) int {
	return num1 / num2
}
