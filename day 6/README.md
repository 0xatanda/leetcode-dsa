# Zigzag Conversion

## [Question](https://leetcode.com/problems/zigzag-conversion/)

## Approach
Approach

Instead of building a visual zigzag grid, I simulate the process by tracking the current row and direction.

Key ideas:

- Each character belongs to a specific row.

- I move downward until the last row is reached.

- Then I reverse direction and move upward.

- Repeat until all characters are processed.

- This keeps the logic simple and avoids unnecessary complexity.


## Algorithm

- If numRows == 1, return the string directly.

- Create a slice of strings — one for each row.

- Initialize:

    - currentRow = 0

    - goingDown = false

- For each character in the string:

    - Append it to rows[currentRow]

    - If at the first or last row, reverse direction

    - Move currentRow up or down based on direction

- Join all rows together and return the result.


## Complexity

- Time: O(n)

- Space: O(n)


``` go
package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)

	currentRow := 0
	goingDown := false

	for _, ch := range s {
		rows[currentRow] += string(ch)

		if currentRow == 0 {
			goingDown = true
		}

		if currentRow == numRows-1 {
			goingDown = false
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}
	return strings.Join(rows, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}

```