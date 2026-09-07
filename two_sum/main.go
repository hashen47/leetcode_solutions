package main

import (
	"log"
	"os"
	"time"
)

type Testcase struct {
	nums   []int
	target int
	want   []int
}

type SolutionFunc func(nums []int, target int) []int

func main() {
	solutionFuncs := []SolutionFunc{twoSum1, twoSum2, twoSum3}

	for _, solutionFunc := range solutionFuncs {
		testcases := []Testcase{
			{
				nums:   []int{2, 7, 11, 15},
				target: 9,
				want:   []int{0, 1},
			},
			{
				nums:   []int{3, 2, 4},
				target: 6,
				want:   []int{1, 2},
			},
			{
				nums:   []int{3, 3},
				target: 6,
				want:   []int{0, 1},
			},
		}

		startTime := time.Now()
		for _, testcase := range testcases {
			got := solutionFunc(testcase.nums, testcase.target)
			if got[0] != testcase.want[0] || got[1] != testcase.want[1] {
				log.Printf("nums: %v, target: %d, want: %v, got: %v\n", testcase.nums, testcase.target, testcase.want, got)
				os.Exit(1)
			}
		}

		log.Printf("ALL TESTCASES ARE PASSED, execution time: %v", time.Since(startTime))
	}
}

func twoSum1(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, n := range nums {
		if i1, ok := m[n]; ok {
			if n+n == target {
				return []int{i1, i}
			}
		} else {
			m[n] = i
		}
	}

	for n, i1 := range m {
		if i2, ok := m[target-n]; ok {
			if i1 == i2 {
				continue
			}
			return []int{i1, i2}
		}
	}

	return []int{0, 0}
}

func twoSum3(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, n := range nums {
		if i1, ok := m[n]; ok {
			if n+n == target {
				return []int{i1, i}
			}
		} else {
			m[n] = i
			if i1, ok := m[target-n]; ok && i1 != i {
				return []int{i1, i}
			}
		}
	}
	return []int{0, 0}
}
