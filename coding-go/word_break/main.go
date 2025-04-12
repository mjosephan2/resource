package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/word-break/?envType=study-plan-v2&envId=top-interview-150
func wordBreak(s string, wordDict []string) bool {
	// DP[len(s)] where DP[0] is true
	// DP[n] that has bool as the value
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if dp[i] {
				break
			}
			wordLength := len(word)
			if wordLength > i {
				continue
			}
			if strings.HasSuffix(s[:i], word) {
				dp[i] = dp[i-wordLength]
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Print(wordBreak("dogs", []string{"dog", "s", "gs"}))
}
