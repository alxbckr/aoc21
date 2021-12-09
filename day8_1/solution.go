package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Line struct {
	formats []string
	digits  []string
}

func main() {
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

	count := 0
	for _, line := range lines {
		for _, digit := range line.digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				count++
			}
		}
	}
	fmt.Printf("Count digits: %v\n", count)
}
