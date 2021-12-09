package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Line struct {
	formats []string
	digits  []string
}

func correctFormat(formats []string, digits []string) []string {
	digitOne := ""
	digitFour := ""
	freqMap := make(map[string]int)
	for _, format := range formats {
		digits := strings.SplitAfter(format, "")
		for _, digit := range digits {
			freqMap[digit] += 1
		}
		if len(format) == 2 {
			digitOne = format
		}
		if len(format) == 4 {
			digitFour = format
		}
	}

	mapping := make(map[string]string)
	for digit, freq := range freqMap {
		switch freq {
		case 6:
			mapping[digit] = "b"
		case 9:
			mapping[digit] = "f"
		case 4:
			mapping[digit] = "e"
		case 8:
			if strings.Contains(digitOne, digit) {
				mapping[digit] = "c"
			} else {
				mapping[digit] = "a"
			}
		case 7:
			if strings.Contains(digitFour, digit) {
				mapping[digit] = "d"
			} else {
				mapping[digit] = "g"
			}
		}
	}

	resultDigits := []string{}
	for _, format := range digits {
		digits := strings.SplitAfter(format, "")
		str := []string{}
		for _, digit := range digits {
			str = append(str, mapping[digit])
		}
		sort.Strings(str)
		resultStr := ""
		for idx := range str {
			resultStr += str[idx]
		}
		resultDigits = append(resultDigits, resultStr)
	}

	return resultDigits
}

func main() {

	digitsTempl := []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}

	lines := []Line{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		inputSpl := strings.Split(input, "|")
		formatsAll := strings.TrimSpace(inputSpl[0])
		digitsAll := strings.TrimSpace(inputSpl[1])
		fmt.Printf("Formats all: %v Digits all %v\n", formatsAll, digitsAll)
		var line Line
		line.formats = strings.Split(formatsAll, " ")
		line.digits = strings.Split(digitsAll, " ")
		lines = append(lines, line)
	}

	results := []int{}
	for _, line := range lines {
		digits := correctFormat(line.formats, line.digits)
		decDigits := []int{}
		for _, digit := range digits {
			for idx, templ := range digitsTempl {
				if digit == templ {
					decDigits = append(decDigits, idx)
				}
			}
		}

		digitRes := 0
		for i := 0; i < len(decDigits); i++ {
			digitRes += decDigits[i] * int(math.Pow10(len(decDigits)-i-1))
		}
		results = append(results, digitRes)
	}

	sum := 0
	for _, res := range results {
		sum += res
	}
	fmt.Printf("Result: %v\n", sum)
}
