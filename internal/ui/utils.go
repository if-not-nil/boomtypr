package ui

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func GetTermDimensions() (int, int, error) {
	fd := int(os.Stdout.Fd())

	width, height, err := term.GetSize(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return 0, 0, err
	}

	return width, height, nil
}

func GetPaddingToCenterVertically(height, lines, padding int) string {
	var b strings.Builder
	halfHeight := height / 2
	halfPadding := padding / 2
	if height%2 == 0 {
		l := (lines + 1) / 2
		padLength := halfHeight - l - halfPadding
		for range padLength {
			b.WriteString("\n")
		}
	} else {
		l := lines / 2
		padLength := halfHeight - l - halfPadding
		for range padLength {
			b.WriteString("\n")
		}
	}

	return b.String()
}
