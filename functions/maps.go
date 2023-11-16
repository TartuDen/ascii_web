package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mapps struct {
	lineNumMap   (map[string]int)
	charCountMap (map[string]int)
	totalMapVert (map[string][]string)
	totalMapHor  (map[string][]string)
}

// CreateLineNumMap reads a file and creates two maps:
// 1. A map of strings to their corresponding line numbers.
// 2. A map of strings to their corresponding line's character count.
func CreateLineNumMap(fileName string) Mapps {
	if fileName == "" {
		fileName = "standard.txt"
	}

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR opening file:", err)
	}
	defer file.Close()

	// Initialize a scanner to read from the file
	scanner := bufio.NewScanner(file)
	lines := []string{}

	// Read each line of the file and store in a slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Create maps to store the results

	mapp := Mapps{
		lineNumMap:   make(map[string]int),
		charCountMap: make(map[string]int),
		totalMapVert: make(map[string][]string),
		totalMapHor:  make(map[string][]string),
	}

	runeCounter := 0

	// Iterate through each line and process
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, " ") { // Checking if the line contains a space character
			mapp.charCountMap[string(32+runeCounter)] = len(line)
			mapp.lineNumMap[string(32+runeCounter)] = i + 1 // Using ASCII code to convert int to string

			mapp.totalMapVert[string(32+runeCounter)] = createTotalMapVert(i, lines)
			mapp.totalMapHor[string(32+runeCounter)] = createTotalMapHorizontal(i, lines)
			runeCounter++
			i += 8 // Skipping the next 8 lines
		}
	}
	// fmt.Println("charCountMap:", charCountMap)

	return mapp
}

func createTotalMapVert(idx int, lines []string) []string {
	tempSlice := []string{}

	for i := 0; i < len(lines[idx]); i++ {
		tempString := ""
		for j := 0; j < 8; j++ {
			tempString = tempString + string(lines[idx+j][i])
		}
		tempSlice = append(tempSlice, tempString)
	}

	return tempSlice

}

func createTotalMapHorizontal(idx int, lines []string) []string {
	tempSlice := []string{}

	for i := 0; i < 8; i++ {
		tempSlice = append(tempSlice, lines[idx+i])
	}
	return tempSlice
}
