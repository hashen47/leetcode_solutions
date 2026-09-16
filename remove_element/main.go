package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(nums[:removeElement(nums, 3)])

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(nums[:removeElement(nums, 2)])
}

func removeElement(nums []int, val int) int {
	length := 0
	for _, n := range nums {
		if val != n {
			nums[length] = n
			length++
		}
	}
	return length
}
