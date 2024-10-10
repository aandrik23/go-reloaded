package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input file> <output file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return

	}

	modifiedText := modifyText(string(content))

	err = os.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Println("Error writing to file", err)
	}
}

func modifyText(text string) string {
	text = handleHexAndBin(text)
	return text
}

func handleHexAndBin(text string) string {
	hexPattern := `(\b[0-9A-Fa-f]+\b) \(hex\)`
	binPattern := `(\b[0-1]+\b) \(bin\)`

	text = regexpReplace(text, hexPattern, func(match string) string {
		hexVal, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(match, " ")[0]), 16, 0)
		return strconv.Itoa(int(hexVal))
	})

	text = regexpReplace(text, binPattern, func(match string) string {
		binVal, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(match, " ")[0]), 2, 0)
		return strconv.Itoa(int(binVal))
	})

	return text
}

func handleCaseChanges(text string) string {
	casePatterns := []string{
		`(\b\w+\b) \(up(?:, (\d+))?\)`,
		`(\b\w+\b) \(low(?:, (\d+))?\)`,
		`(\b\w+\b) \(cap(?:, (\d+))?\)`,
	}

// 	for _, pattern := range casePatterns {
// 		text = regexpReplace(text, pattern, func(match string, groups []string) string {
// 			word := groups[0]
// 			count := 1
// 			if len(groups) > 1 && groups[1] != "" {
// 				count, _ = strconv.Atoi(groups[1])
// 			}

// 			switch {
// 			case strings.Contains(pattern, "(up"):
// 				return strings.ToUpper(strings.Join(strings.Fields(text), " ")[:count])
// 			case strings.Contains(pattern, "(low"):
// 				return strings.ToLower(strings.Join(strings.Fields(text), " ")[:count])
// 			case strings.Contains(pattern, "(cap"):
// 				return capitalizeWords(strings.Join(strings.Fields(text), " ")[:count])
// 			}
// 			return match
// 		})
// 	}

// 	return text
// }

func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

func regexpReplace(text string, pattern string, replaceFunc func(string) string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllStringFunc(text, replaceFunc)
}
