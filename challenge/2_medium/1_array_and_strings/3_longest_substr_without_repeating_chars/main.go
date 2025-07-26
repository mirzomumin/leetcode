package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var maxLength int
	var seenChars = make(map[rune]int)
	var left int
	for right, char := range s {
		seenCharIdx, ok := seenChars[char]
		seenChars[char] = right
		if !ok || seenCharIdx < left {
			continue
		}

		length := right - left
		maxLength = max(maxLength, length)
		left = seenCharIdx + 1
	}

	return max(len(s)-left, maxLength)
}

func main() {
	strs := []string{"abcabcbb", "bbbbb", "pwwkew", "aab", "abcdefg"}
	for _, str := range strs {
		length := lengthOfLongestSubstring(str)
		fmt.Printf("LENGTH of longest substring in `%s`: %d\n", str, length)
	}
}
