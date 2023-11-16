// Package functions provides utility functions for colorizing words in a text file.
package functions

import (
	"fmt"
	"os"
	"strings"
)

var TotalWordsToColor int

func writeOutput(str string, outputFile string, file *os.File) {

	_, err2 := file.WriteString(str) //write into a file
	if err2 != nil {
		fmt.Println("ERROR writing into file", err2)
	}
}

// RunFunction processes the inputString, reading from a file specified by toReadFrom,
// and colorizes a specified number of words using the given color.
func RunFunction(inputString, banner, lettersToColor string) []string {
	resultString := ""
	resultSlice := []string{}
	mapp := CreateLineNumMap(banner)
	fSplit := strings.Split(inputString, "\n")
	for _, elem := range fSplit {
		for i := 0; i < 8; i++ {
			for idx := 0; idx < len(elem); idx++ {
				ch := string(elem[idx])

				resultString = resultString + mapp.totalMapHor[ch][i]

			}

			resultSlice = append(resultSlice, resultString)
			resultString = ""

		}
	}
	return resultSlice
}
