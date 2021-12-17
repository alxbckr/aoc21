package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InsertRules map[string]string

func processRules(template string, rules InsertRules) string {
	chrPrev := ""
	builder := strings.Builder{}
	for i := 0; i < len(template); i++ {
		chr := string(template[i])
		elem, exists := rules[chrPrev+chr]
		if exists {
			builder.WriteString(elem + chr)
		} else {
			builder.WriteString(chr)
		}
		chrPrev = chr
	}
	return builder.String()
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer f.Close()

	template := ""
	rules := InsertRules{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		if template == "" {
			template = input
		} else if input == "" {
			continue
		} else {
			seq := ""
			ins := ""
			fmt.Sscanf(input, "%s -> %s", &seq, &ins)
			rules[seq] = ins
		}
	}

	for i := 0; i < 10; i++ {
		template = processRules(template, rules)
		fmt.Printf("Day %v\n", i)
	}

	freq := make(map[string]int)
	for _, chr := range strings.SplitAfter(template, "") {
		freq[chr] += 1
	}

	mostFreqCnt := 0
	mostFreq := ""
	leastFreqCnt := 0
	leastFreq := ""
	for k, v := range freq {
		if v > mostFreqCnt {
			mostFreq = k
			mostFreqCnt = v
		}
		if v < leastFreqCnt || leastFreqCnt == 0 {
			leastFreq = k
			leastFreqCnt = v
		}
	}

	fmt.Printf("Most common %v occurs %v\n", mostFreq, mostFreqCnt)
	fmt.Printf("Least common %v occurs %v\n", leastFreq, leastFreqCnt)
	fmt.Printf("Result %v\n", mostFreqCnt-leastFreqCnt)
}
