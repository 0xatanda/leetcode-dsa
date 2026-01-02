# Add Two Number (Linked List)

## [Question](https://leetcode.com/problems/add-two-numbers/)

## Approach

---

I solve this problem by simulating elementary addition with carry, traversing both linked lists simultaneously.

Since the digits are stored in reverse order, we can add corresponding nodes directly, just like adding numbers from right to left.

At each step:

- Add the values of the current nodes (if present) and the carry

- Store the digit (sum % 10) in a new node

- Carry over (sum / 10) to the next step

The process continues until both lists are fully traversed and no carry remains.

--- 

## Algorithm

- Initialize a dummy head for the result list.

- Set a carry variable to 0.

- While at least one list has nodes or carry > 0:

    - Read values from current nodes (use 0 if a list is exhausted)

    - Compute sum = val1 + val2 + carry

    - Update carry = sum / 10

    - Append a node with value sum % 10 to the result list

- Return dummy.next

## Complexity
    - Time: ```O(max(m, n))```

    - Space: ```O(max(m, n))``` (result list)

Where m and n are the lengths of the two linked lists.

``` python 

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        dummy = ListNode(0)
        current = dummy
        carry = 0

        while l1 or l2 or carry:
            v1 = l1.val if l1 else 0
            v2 = l2.val if l2 else 0

            total = v1 + v2 + carry
            carry = total // 10

            current.next = ListNode(total % 10)
            current = current.next

            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next

        return dummy.next

```