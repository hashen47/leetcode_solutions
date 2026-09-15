package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	end := 0
	minLen := len(strs[0])

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	for i := 0; i < minLen; i++ {
		ch := strs[0][i]
		isAllEqual := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != ch {
				isAllEqual = false
			}
		}
		if !isAllEqual {
			break
		}
		end++
	}

	if end == 0 {
		return ""
	}

	return strs[0][:end]
}
