package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return s[start : start+maxLen]
}

func main() {
	longestPalindrome("babad")
}
