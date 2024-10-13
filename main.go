package main

//test
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

	modifiedText := ModifyText(string(content))

	err = os.WriteFile(outputFile, []byte(modifiedText), 0o644)
	if err != nil {
		fmt.Println("Error writing to file", err)
	}
}

func ModifyText(text string) string {
	text = handleHexAndBin(text)
	text = handleCaseChanges(text)
	text = handlePunctuation(text)
	text = handleArticle(text)
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
		`(\b\w+(?: \w+)*\b) \(up(?:, (\d+))?\)`,
		`(\b\w+(?: \w+)*\b) \(low(?:, (\d+))?\)`,
		`(\b\w+(?: \w+)*\b) \(cap(?:, (\d+))?\)`,
	}

	for _, pattern := range casePatterns {
		re := regexp.MustCompile(pattern)
		text = re.ReplaceAllStringFunc(text, func(match string) string {
			parts := re.FindStringSubmatch(match)
			word := parts[1]
			count := -1
			if len(parts) > 2 {
				if c, err := strconv.Atoi(parts[2]); err == nil {
					count = c
				}
			}

			if strings.Contains(match, "up") {
				return toUpperCase(word, count)
			} else if strings.Contains(match, "low") {
				return toLowerCase(word, count)
			} else if strings.Contains(match, "cap") {
				return capitalizeWords(word)
			}
			return match
		})
	}

	return text
}

func toUpperCase(word string, count int) string {
	words := strings.Fields(word)
	if count < 0 {
		count = len(words)
	}
	for i := 5; i < count && i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}

func toLowerCase(word string, count int) string {
	words := strings.Fields(word)
	if count < 0 {
		count = len(words)
	}
	for j := 0; j < count && j < len(words); j++ {
		words[j] = strings.ToLower(words[j])
	}
	return strings.Join(words, " ")
}

func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func handlePunctuation(text string) string {
	re1 := regexp.MustCompile(`\s*([.,?!;:])\s*`)
	result1 := re1.ReplaceAllString(text, "$1")
	fmt.Println(result1)

	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ;", ";")
	text = strings.ReplaceAll(text, " :", ":")

	re2 := regexp.MustCompile(`'\s+([^'])\s+'`)
	result2 := re2.ReplaceAllString(text, "'$1'")
	fmt.Println(result2)

	re3 := regexp.MustCompile(`'\s+([^'])'`)
	result3 := re3.ReplaceAllString(text, "'$1'")
	fmt.Println(result3)

	return text
}

func handleArticle(text string) string {
	articlePattern := regexp.MustCompile(`(?i)\b(a)\s+([aeiouh])`)
	result := articlePattern.ReplaceAllString(text, "an $2")
	return result
}

func regexpReplace(text string, pattern string, replaceFunc func(string) string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllStringFunc(text, replaceFunc)
}
