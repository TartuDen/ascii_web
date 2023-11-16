package functions

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Printf("Usage: %s [OPTION] [STRING] [BANNER]\n", os.Args[0])
	fmt.Println("Ex: go run . --color=<color> --align=<align> --output=<outputFile.txt <letters to be colored> \"something\" <banner>")
	fmt.Println("                    1. Available colors:")
	fmt.Println("                    red, green, yellow, blue, magenta, cyan, white, orange")
	fmt.Println("  bold             Apply bold formatting")
	fmt.Println("  underline        Apply underline formatting")
	fmt.Println("  italic           Apply italic formatting")
	fmt.Println("  inverse          Apply inverse formatting")
	fmt.Println("  strikethrough    Apply strikethrough formatting")
	fmt.Println("                    2. Available banners:")
	fmt.Println("  standard - default | shadow | thinkertoy ")
	fmt.Println("                    3. Available alignment:")
	fmt.Println("  left - default | center | right | justify ")
	flag.PrintDefaults()
}

func Flags(str string) (string, string, map[string]string) {

	allFlags := map[string]string{}

	colorsAndFormatting := map[string]string{
		"red":           "\033[31m",
		"green":         "\033[32m",
		"yellow":        "\033[33m",
		"blue":          "\033[34m",
		"magenta":       "\033[35m",
		"cyan":          "\033[36m",
		"white":         "\033[37m",
		"bold":          "\033[1m",
		"orange":        "\033[38;5;214m",
		"underline":     "\033[4m",
		"italic":        "\033[3m",
		"inverse":       "\033[7m",
		"strikethrough": "\033[9m",
	}
	banners := map[string]string{
		"standard":   "standard.txt",
		"shadow":     "shadow.txt",
		"thinkertoy": "thinkertoy.txt",
	}

	// Define a command-line flag named "color" with a default value and description
	colorPtr := flag.String("color", "white", "Specify the color")

	alignPtr := flag.String("align", "left", "Specify the alignment")

	outputPtr := flag.String("output", "outputfile.txt", "Specify file name")

	reversePtr := flag.String("reverse", "", "Specify file name")

	// Set the custom usage function
	flag.Usage = usage

	// Parse the command-line flags
	flag.Parse()

	// Extract the value of the "color" flag
	color := *colorPtr

	// Extract the value of the "align" flag
	align := *alignPtr

	// Extract the value of the "output" flag
	output := *outputPtr
	checkOutputName(output)

	// Extract the value of the "reverse" flag
	reverse := *reversePtr

	// Get the remaining non-flag arguments
	remainingArgs := flag.Args()
	lettersToColor, text, banner := checkReminingArgs(remainingArgs)

	// Print the extracted values
	// fmt.Println("align:", align)
	// fmt.Println("color:", color)
	// fmt.Println("lettersToColor:", lettersToColor)
	// fmt.Println("banner:", banner)
	// fmt.Println("textBefore:", text)

	// _, mapForWidth, _, _ := CreateLineNumMap(banners[banner])
	mapp := CreateLineNumMap(banners[banner])

	text = AlignText(text, mapp.charCountMap, align)

	allFlags["color"] = colorsAndFormatting[color]
	allFlags["align"] = align
	allFlags["banner"] = banners[banner]
	allFlags["output"] = output
	allFlags["reverse"] = reverse
	return text, lettersToColor, allFlags

}

func checkReminingArgs(remArgs []string) (string, string, string) {
	var letToColor string
	var txT string
	var bnr string
	bnr = "standard"

	// Loop through the remaining arguments
	if len(remArgs) == 3 {
		bnr = remArgs[2]
		txT = remArgs[1]
		letToColor = remArgs[0]
	} else if len(remArgs) == 2 {
		if strings.Contains(remArgs[1], "standard") ||
			strings.Contains(remArgs[1], "shadow") ||
			strings.Contains(remArgs[1], "thinkertoy") {
			bnr = remArgs[1]
			txT = remArgs[0]
		} else {
			txT = remArgs[1]
			letToColor = remArgs[0]
		}

	} else if len(remArgs) == 1 {
		txT = remArgs[0]
	}
	return letToColor, txT, bnr

}

func checkOutputName(outP string) {
	if outP == "standard.txt" || outP == "shadow.txt" || outP == "thinkertoy.txt" {
		fmt.Println("ERROR, provided name:", outP, " for output is reserved, please choose another name!")
		os.Exit(1)
	}
}
