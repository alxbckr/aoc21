package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InsertRules map[string]string
type Freq map[string]int

func processRules(templatePairs Freq, rules InsertRules, freq Freq) Freq {
	result := make(map[string]int)
	for k, v := range rules {
		templ1 := string(k[0]) + v
		templ2 := v + string(k[1])
		val, exists := templatePairs[k]
		if exists {
			result[templ1] += val
			result[templ2] += val
			freq[v] += val
		}
	}
	return result
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

	templatePairs := make(Freq)
	chrPrev := ""
	for i := 0; i < len(template); i++ {
		chr := string(template[i])
		pair := chr
		if chrPrev != "" {
			pair = chrPrev + pair
			templatePairs[pair] += 1
		}
		chrPrev = chr
	}

	freq := make(Freq)
	for _, chr := range strings.SplitAfter(template, "") {
		freq[chr] += 1
	}

	for i := 0; i < 40; i++ {
		templatePairs = processRules(templatePairs, rules, freq)
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
