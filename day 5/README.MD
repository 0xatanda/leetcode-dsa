# Longest Palindromic Substring 

## [Question](https://leetcode.com/problems/longest-palindromic-substring/)

## Approach 
This problem is efficiently solved using the expand-around-center technique.

A palindrome mirrors around its center. Each character (and each gap between characters) can act as a potential center. From each center, we expand outward while the characters on both sides are equal.

Two cases are considered:

- Odd-length palindromes (single center character)

- Even-length palindromes (center between two characters)

The longest palindrome found during these expansions is returned.

## Algorithm

- If the string length is less than 2, return the string.

- For each index in the string:

    - Expand around the center (i, i) for odd-length palindromes.

    - Expand around the center (i, i+1) for even-length palindromes.

- Track the start index and maximum length of the longest palindrome found.

- Return the substring using the tracked indices.

## Complexity

- Time: O(n²)

- Space: O(1)


``` go 
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
```