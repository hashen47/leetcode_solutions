package main

import "log"

type Testcase struct {
	arr  []int
	k    int
	want int
}

func main() {
	testcases := []Testcase{
		{
			arr:  []int{5, 2, -1, 0, 3},
			k:    3,
			want: 6,
		},
		{
			arr:  []int{1, 4, 2, 10, 23, 3, 1, 0, 20},
			k:    4,
			want: 39,
		},
	}

	isAllTestcasesArePassed := true
	for _, tc := range testcases {
		got := maxSum(tc.arr, tc.k)
		if got != tc.want {
			isAllTestcasesArePassed = false
			log.Printf("arr: %v, got: %d, want: %d\n", tc.arr, got, tc.want)
		}
	}

	if isAllTestcasesArePassed {
		log.Println("ALL TESTCASES ARE PASSED")
	}
}

func maxSum(arr []int, k int) int {
	maxSum := 0
	sum := 0

	for i := range k {
		sum += arr[i]
	}

	maxSum = sum

	for i := k; i < len(arr); i++ {
		sum += arr[i]
		sum -= arr[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
