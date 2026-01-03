# Longest Substring Without Repeating Characters

[Questions](https://leetcode.com/problems/longest-substring-without-repeating-characters/)


## Approach 
This problem is best solved using the sliding window technique with a hash map (or set) to track characters within the current window.

Two pointers (left and right) define a window that expands as we iterate through the string. When a duplicate character is encountered, the window is adjusted by moving the left pointer forward until the duplicate is removed.

By maintaining the window in this way, we ensure each character is processed at most twice.

## Algorithm
1. Initialize a hash map to store characters and their most recent indices.

2. Set left = 0 and maxLength = 0.

3. Iterate through the string with index right:

    - If the current character exists in the map and its index is ≥ left, move left to one position after the previous occurrence.

    - Update the character’s index in the map.

    - Update maxLength as right - left + 1.

4. Return maxLength.

## Complexity
Time: O(n)

Space: O(min(n, m)), where m is the size of the character set


``` python 

class Solution(object):
    def lengthOfLongestSubstring(self, s):
        charIndex = {}
        left = 0
        maxLength = 0

        for right, ch in enumerate(s):
            if ch in charIndex and charIndex[ch] >= left:
                left = charIndex[ch] + 1

            charIndex[ch] = right
            maxLength = max(maxLength, right - left + 1)

        return maxLength
```