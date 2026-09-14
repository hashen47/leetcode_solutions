package main

import "fmt"

var RomanIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

type Testcase struct {
	input string
	want  int
}

func main() {
	testcases := []Testcase{
		{
			input: "III",
			want:  3,
		},
		{
			input: "LVIII",
			want:  58,
		},
		{
			input: "MCMXCIV",
			want:  1994,
		},
	}

	isAllPassed := true
	for _, tc := range testcases {
		got := romanToInt(tc.input)
		if got != tc.want {
			isAllPassed = false
			fmt.Printf("roman: %s, want: %d, got: %d\n", tc.input, tc.want, got)
		}
	}

	if isAllPassed {
		fmt.Println("ALL TESTCASES ARE PASSED...")
	}
}

func romanToInt(s string) int {
	number := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			number += RomanIntMap['M']
		case 'D':
			number += RomanIntMap['D']
		case 'C':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'D':
					number += RomanIntMap['D'] - RomanIntMap['C']
					i++
				case 'M':
					number += RomanIntMap['M'] - RomanIntMap['C']
					i++
				default:
					number += RomanIntMap['C']
				}
			} else {
				number += RomanIntMap['C']
			}
		case 'L':
			number += RomanIntMap['L']
		case 'X':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'L':
					number += RomanIntMap['L'] - RomanIntMap['X']
					i++
				case 'C':
					number += RomanIntMap['C'] - RomanIntMap['X']
					i++
				default:
					number += RomanIntMap['X']
				}
			} else {
				number += RomanIntMap['X']
			}
		case 'V':
			number += RomanIntMap['V']
		case 'I':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'V':
					number += RomanIntMap['V'] - RomanIntMap['I']
					i++
				case 'X':
					number += RomanIntMap['X'] - RomanIntMap['I']
					i++
				default:
					number += RomanIntMap['I']
				}
			} else {
				number += RomanIntMap['I']
			}
		}
	}

	return number
}
