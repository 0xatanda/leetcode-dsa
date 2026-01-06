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
