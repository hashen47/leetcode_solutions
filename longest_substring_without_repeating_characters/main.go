package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type SolutionFunc func(s string) int

type Testcase struct {
	text string
	want int
}

func main() {
	testcases := []Testcase{
		{
			text: "abcabcbb",
			want: 3,
		},
		{
			text: "bbbbb",
			want: 1,
		},
		{
			text: "pwwkew",
			want: 3,
		},
		{
			text: "S",
			want: 1,
		},
		{
			text: "abcd",
			want: 4,
		},
	}

	solutionFuncs := []SolutionFunc{
		lengthOfLongestSubstring1,
		lengthOfLongestSubstring2,
	}

	for _, fn := range solutionFuncs {
		startTime := time.Now()
		for _, tc := range testcases {
			got := fn(tc.text)
			if got != tc.want {
				log.Printf("text: %q, want: %d, got: %d\n", tc.text, tc.want, got)
				os.Exit(1)
			}
		}
		log.Printf("[ALL TESTCASES ARE PASSED] elapsed time: %v\n", time.Since(startTime))
	}

}

func lengthOfLongestSubstring1(s string) int {
	maxLength := 0

	m := make(map[byte]int, 0)
	left := 0
	length := 0
	for right := 0; right < len(s); right++ {
		if i, ok := m[s[right]]; ok {
			if maxLength < length {
				maxLength = length
			}
			for j := left; j <= i; j++ {
				delete(m, s[j])
				length--
			}
			m[s[right]] = right
			left = i + 1
		}
		m[s[right]] = right
		length++
	}

	if maxLength < length {
		return length
	}

	return maxLength
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLen := 0
	left := 0
	charMap := make(map[byte]int, 0)
	right := 0
	for right = 0; right < len(s); right++ {
		ch := s[right]

		if i, ok := charMap[ch]; ok && i >= left {
			if maxLen < right-left {
				maxLen = right - left
			}
			left = i + 1
		}

		charMap[ch] = right
	}

	if right-left > maxLen {
		return right - left
	}

	return maxLen
}

func printText(text string, start, end int) {
	for ; start <= end; start++ {
		fmt.Printf("%c", text[start])
	}
	fmt.Println("")
}
