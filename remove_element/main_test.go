package main

import (
	"testing"
)

type Testcase struct {
	nums    []int
	val     int
	wantLen int
	want    []int
}

func TestRemoveElement(t *testing.T) {
	testcases := []Testcase{
		{
			nums:    []int{3, 2, 2, 3},
			val:     3,
			wantLen: 2,
			want:    []int{2, 2},
		},
		{
			nums:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:     2,
			wantLen: 5,
			want:    []int{0, 1, 3, 0, 4},
		},
	}

	for _, tc := range testcases {
		duplicate := append([]int(nil), tc.nums...)
		gotLen := removeElement(duplicate, tc.val)
		assertNums(t, tc.nums, tc.wantLen, tc.want, gotLen, duplicate[:gotLen])
	}
}

func assertNums(t testing.TB, initial []int, wantLen int, want []int, gotLen int, got []int) {
	t.Helper()

	if gotLen != wantLen {
		t.Errorf("initial: %v, wantLen: %d, gotLen: %d\n", initial, wantLen, gotLen)
		return
	}

	for i := range wantLen {
		if got[i] != want[i] {
			t.Errorf("initial: %v, want: %v, got: %v\n", initial, want, got)
			return
		}
	}
}
