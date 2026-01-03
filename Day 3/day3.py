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