# Median of Two Sorted Arrays

## [Question](https://leetcode.com/problems/median-of-two-sorted-arrays/)

---

## Approach

To achieve logarithmic time complexity, this problem is solved using **binary search on the smaller array**.

The idea is to partition both arrays such that:

* The left half contains exactly half of the total elements.
* Every element in the left half is less than or equal to every element in the right half.

By adjusting the partition using binary search, we find the correct split where:

```
max(left partitions) ≤ min(right partitions)
```

Once the correct partition is found:

* If the total number of elements is odd, the median is the maximum of the left partition.
* If even, the median is the average of the maximum of the left partition and the minimum of the right partition.

---

## Algorithm

1. Ensure `nums1` is the smaller array.
2. Perform binary search on `nums1`.
3. For each partition:

   * Compute corresponding partition in `nums2`.
   * Check partition validity.
4. Adjust search space until a valid partition is found.
5. Compute and return the median.

---

## Complexity

* **Time:** `O(log(min(m, n)))`
* **Space:** `O(1)`

---

## Python Implementation

```python
class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1

        m, n = len(nums1), len(nums2)
        left, right = 0, m

        while left <= right:
            partitionX = (left + right) // 2
            partitionY = (m + n + 1) // 2 - partitionX

            maxLeftX = float('-inf') if partitionX == 0 else nums1[partitionX - 1]
            minRightX = float('inf') if partitionX == m else nums1[partitionX]

            maxLeftY = float('-inf') if partitionY == 0 else nums2[partitionY - 1]
            minRightY = float('inf') if partitionY == n else nums2[partitionY]

            if maxLeftX <= minRightY and maxLeftY <= minRightX:
                if (m + n) % 2 == 0:
                    return (max(maxLeftX, maxLeftY) + min(minRightX, minRightY)) / 2.0
                else:
                    return max(maxLeftX, maxLeftY)
            elif maxLeftX > minRightY:
                right = partitionX - 1
            else:
                left = partitionX + 1
```
