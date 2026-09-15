package main

import (
	"strings"
	"testing"
)

type Testcase struct {
	strs []string
	want string
}

func TestLongestCommonPrefix(t *testing.T) {
	testcases := []Testcase{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "a",
		},
	}

	for _, tc := range testcases {
		got := longestCommonPrefix(tc.strs)
		assertString(t, got, tc.want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if strings.Compare(got, want) != 0 {
		t.Errorf("expected to be equal, but got: %q, want: %q\n", got, want)
	}
}
