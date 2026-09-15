package main

import "testing"

type Testcase struct {
	nums []int
	want []int
}

func TestRemoveDuplicates(t *testing.T) {
	testcases := []Testcase{
		{
			nums: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range testcases {
		gotLen := removeDuplicates(tc.nums)
		assertArrayEqual(t, gotLen, tc.nums, tc.want)
	}
}

func assertArrayEqual(t testing.TB, gotLen int, initial, want []int) {
	t.Helper()

	if gotLen != len(want) {
		t.Errorf("initial array: %v, want: %v, got: %v\n", initial, want, initial[:len(want)])
		return
	}

	for i := range len(want) {
		if initial[i] != want[i] {
			t.Errorf("initial array: %v, want: %v, got: %v\n", initial, want, initial[:len(want)])
			return
		}
	}
}
