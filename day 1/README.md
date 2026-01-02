# Two Sum

## [Question ](https://leetcode.com/problems/two-sum/)


Given an array of integers `nums` and an integer `target`, return the **indices of the two numbers** such that they add up to `target`.
You may assume that **each input has exactly one solution** and the same element **cannot be used twice**.

---

## Approach

I solve this using a **single-pass hash map**. The key idea is to track numbers we have already seen as we iterate through the array.

For each number `nums[i]`, I compute its complement with respect to the target:

```
complement = target - nums[i]
```

If that complement has already been seen (and stored in the map), then I’ve found the pair and can return their indices. If not, I record the current number with its index in the map. This ensures we only scan the list once.

This approach gives:

* **Time Complexity:** **O(n)** — one pass over the list with constant-time lookups
* **Space Complexity:** **O(n)** — storage used by the map

---

## Implementations

### Python

```python
def twoSum(nums, target):
    prevMap = {}
    for i, n in enumerate(nums):
        diff = target - n
        if diff in prevMap:
            return [prevMap[diff], i]
        prevMap[n] = i
```

---

### Java

```java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int diff = target - nums[i];
            if (map.containsKey(diff)) {
                return new int[] { map.get(diff), i };
            }
            map.put(nums[i], i);
        }
        return new int[] {};
    }
}
```

---

### Go

```go
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, n := range nums {
        diff := target - n
        if idx, ok := m[diff]; ok {
            return []int{idx, i}
        }
        m[n] = i
    }
    return nil
}
```

---

### TypeScript

```ts
function twoSum(nums: number[], target: number): number[] {
    const map = new Map<number, number>();

    for (let i = 0; i < nums.length; i++) {
        const diff = target - nums[i];
        if (map.has(diff)) {
            return [map.get(diff)!, i];
        }
        map.set(nums[i], i);
    }
    return [];
}
```

---

## Complexities

* **Time:** `O(n)` — single pass with hash lookups
* **Space:** `O(n)` — storing numbers in the map
