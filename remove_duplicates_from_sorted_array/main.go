package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	uniqCount := 0
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		if n != nums[i] {
			nums[uniqCount] = n
			n = nums[i]
			uniqCount++
		}
	}
	nums[uniqCount] = n
	uniqCount++
	return uniqCount
}
