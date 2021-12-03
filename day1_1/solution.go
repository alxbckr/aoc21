package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	previousMeasure := 0
	result := 0
	var currentMeasure int
	for scanner.Scan() {
		var err error
		if currentMeasure, err = strconv.Atoi(scanner.Text()); err != nil {
			break
		}
		if previousMeasure != 0 && currentMeasure > previousMeasure {
			result += 1
		}
		previousMeasure = currentMeasure
	}
	fmt.Println(result)
}
