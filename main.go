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

// handleHexAndBin converts hexadecimal and binary values in the text to their decimal equivalents.
func handleHexAndBin(text string) string {
	hexPattern := `(\b[0-9A-Fa-f]+\b) \(hex\)` //matches one or more characters that are digits (0-9) or letters (A-F or a-f).
	binPattern := `(\b[0-1]+\b) \(bin\)`       //matches one or more characters that are either 0 or 1.

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

// upPattern, lowPattern, and capPattern handle word transformations specified as '(up)', '(low)', or '(cap)'
func handleCaseChanges(text string) string {
	upPattern := regexp.MustCompile(`([A-Za-z]+) \(up\)`)
	lowPattern := regexp.MustCompile(`([A-Za-z]+) \(low\)`)
	capPattern := regexp.MustCompile(`([A-Za-z]+) \(cap\)`)
	countedUpLowCapPattern := regexp.MustCompile(`([A-Za-z\s]+) \((up|low|cap), (\d+)\)`)

	text = upPattern.ReplaceAllStringFunc(text, func(match string) string {
		word := upPattern.FindStringSubmatch(match)[1]
		return strings.ToUpper(word)
	})

	text = lowPattern.ReplaceAllStringFunc(text, func(match string) string {
		word := lowPattern.FindStringSubmatch(match)[1]
		return strings.ToLower(word)
	})

	text = capPattern.ReplaceAllStringFunc(text, func(match string) string {
		word := capPattern.FindStringSubmatch(match)[1]
		return capitalize(word)
	})

	// countedUpLowCapPattern handles transformations where the case change is limited to a specified number
	text = countedUpLowCapPattern.ReplaceAllStringFunc(text, func(match string) string {
		words := strings.Fields(countedUpLowCapPattern.FindStringSubmatch(match)[1])
		action := countedUpLowCapPattern.FindStringSubmatch(match)[2]
		count, _ := strconv.Atoi(countedUpLowCapPattern.FindStringSubmatch(match)[3])
		switch action {
		case "up":
			for i := len(words) - 1; i >= len(words)-count && i >= 0; i-- {
				words[i] = strings.ToUpper(words[i])
			}
		case "low":
			for i := len(words) - 1; i >= len(words)-count && i >= 0; i-- {
				words[i] = strings.ToLower(words[i])
			}
		case "cap":
			for i := len(words) - 1; i >= len(words)-count && i >= 0; i-- {
				words[i] = capitalize(words[i])
			}
		}
		return strings.Join(words, " ")
	})
	return text
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + string(word[1:])
}

func handlePunctuation(text string) string {
	punctuation := regexp.MustCompile(`(\S)\s*([.,!?;:])`)
	text = punctuation.ReplaceAllString(text, `$1$2$3`)

	punctuationCluster := regexp.MustCompile(`([.,!?;:])\s*([.,!?;:])`)
	text = punctuationCluster.ReplaceAllString(text, `$1$2`)

	punctuationSpace := regexp.MustCompile(`([.,!?;:])([^\s.,!?;:'"\"])`)
	text = punctuationSpace.ReplaceAllString(text, `$1 $2`)

	quotation := regexp.MustCompile(`\s*'(\w+)'(\s*)`)
	text = quotation.ReplaceAllString(text, `'$1'`)

	quotationMult := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = quotationMult.ReplaceAllString(text, `'$1'`)

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
