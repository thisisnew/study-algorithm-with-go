package main

import "fmt"

type Students struct {
	num     int
	lost    bool
	reserve bool
}

func main() {

	n := 5
	nums := []int{2, 4}
	reserve := []int{2, 4}

	fmt.Println(체육복(n, nums, reserve))
}

func 체육복(n int, lost []int, reserve []int) int {

	students := generateStudents(n)

	for i := 0; i < len(students); i++ {
		students[i].lost = isLostStudent(lost, students[i].num)
		students[i].reserve = isReservedStudent(reserve, students[i].num)
	}

	return countStudentWhoHasGymSuit(students)
}

func generateStudents(n int) []Students {
	var result []Students

	for i := 0; i < n; i++ {

		r := Students{
			num: i + 1,
		}

		result = append(result, r)
	}

	return result
}

func isLostStudent(lost []int, n int) bool {

	for _, ln := range lost {
		if ln == n {
			return true
		}
	}

	return false
}

func isReservedStudent(reserve []int, n int) bool {

	for _, rn := range reserve {
		if rn == n {
			return true
		}
	}

	return false
}

func countStudentWhoHasGymSuit(students []Students) int {

	var result int

	for i := 0; i < len(students); i++ {

		if !students[i].lost {
			continue
		}

		if students[i].lost && students[i].reserve {
			students[i].reserve = false
			students[i].lost = false
			continue
		}

		if students[i].lost && i >= 0 && i < len(students)-1 {
			if students[i+1].reserve {
				students[i+1].reserve = false
				students[i].lost = false
				continue
			}
		}

		if students[i].lost && i > 0 && i <= len(students)-1 {
			if students[i-1].reserve {
				students[i-1].reserve = false
				students[i].lost = false
				continue
			}
		}

	}

	for _, st := range students {
		if !st.lost {
			result++
		}
	}

	return result
}
