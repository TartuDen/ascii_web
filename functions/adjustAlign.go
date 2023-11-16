package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// AlignText aligns text based on the specified alignment (right, center, justify)
// within the current terminal window's width.
//
// Parameters:
//
//	str    : The input text to align.
//	wMap   : A map that assigns widths to characters.
//	align  : The alignment type ("right", "center", or "justify").
//
// Returns:
//
//	string : The aligned text as a single string.
func AlignText(str string, wMap map[string]int, align string) string {
	// Get the current terminal window's width
	width, _, _ := getTerminalSize()

	// Split the input text into lines using "\n" as the separator
	newSlice := strings.Split(str, "\\n")

	if align == "right" {
		// Align text to the right
		for idx, elem := range newSlice {
			tempStr := ""
			lenOfLine := 0
			totalSpaces := 0

			// Calculate the total width of the line
			for _, ch := range elem {
				lenOfLine += wMap[string(ch)]
			}

			// Calculate the total number of spaces needed for alignment
			totalSpaces = width - lenOfLine
			totalSpaces = (totalSpaces) / wMap[" "]

			// Add the appropriate number of spaces to align the text to the right
			for i := 1; i < totalSpaces; i++ {
				tempStr = tempStr + " "
			}
			tempStr = tempStr + elem

			newSlice[idx] = tempStr
		}

	} else if align == "center" {
		// Center-align text
		for idx, elem := range newSlice {
			tempStr := ""
			lenOfLine := 0
			totalSpacesOneSide := 0

			// Calculate the total width of the line
			for _, ch := range elem {
				lenOfLine += wMap[string(ch)]
			}

			// Calculate the total number of spaces needed on one side for center alignment
			totalSpacesOneSide = width/2 - lenOfLine/2
			totalSpacesOneSide = (totalSpacesOneSide) / wMap[" "]

			// Add the appropriate number of spaces on both sides to center-align the text
			for i := 1; i <= totalSpacesOneSide; i++ {
				tempStr = tempStr + " "
			}
			tempStr = tempStr + elem + tempStr

			newSlice[idx] = tempStr
		}

	} else if align == "justify" {
		// Justify-align text
		for idx, elem := range newSlice {
			innerSlice := strings.Split(elem, " ")
			wNum := 2
			if strings.Contains(elem, " ") {
				wNum = len(innerSlice)
			}

			tempStr := ""
			lenOfLine := 0
			totalSpace := 0

			// Calculate the total width of the line (excluding spaces)
			for _, ch := range strings.Join(innerSlice, "") {
				lenOfLine += wMap[string(ch)]
			}

			// Calculate the total number of spaces needed for justification
			totalSpace = width/(wNum-1) - lenOfLine/(wNum-1)
			totalSpace = (totalSpace) / wMap[" "]

			// Add the appropriate number of spaces between words to justify-align the text
			for i := 0; i < totalSpace; i++ {
				tempStr = tempStr + " "
			}
			tempStr = strings.Join(innerSlice, tempStr)

			newSlice[idx] = tempStr
		}
	}

	// Join the lines with "\\n" to form the aligned text
	return strings.Join(newSlice, "\\n")
}

// getTerminalSize retrieves the current terminal window's size (width and height).
//
// Returns:
//
//	int    : The width of the terminal window.
//	int    : The height of the terminal window.
//	error  : An error, if any, that occurred during retrieval.
func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	sizeStr := strings.TrimSpace(string(out))
	sizeParts := strings.Split(sizeStr, " ")
	if len(sizeParts) != 2 {
		return 0, 0, fmt.Errorf("unexpected output from stty: %s", sizeStr)
	}

	width, err := strconv.Atoi(sizeParts[1])
	if err != nil {
		return 0, 0, err
	}

	height, err := strconv.Atoi(sizeParts[0])
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}
