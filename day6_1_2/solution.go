package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const numDays int = 256

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	input := ""
	for scanner.Scan() {
		input = scanner.Text()
		break
	}

	fishes := make([]int, 9)
	for _, str := range strings.Split(input, ",") {
		fishAge, _ := strconv.Atoi(str)
		fishes[fishAge]++
	}

	fmt.Printf("Initial Fishes %v\n", fishes)

	for day := 0; day < numDays; day++ {
		fishes0 := fishes[0]
		for idx := 0; idx < 8; idx++ {
			fishes[idx] = fishes[idx+1]
		}
		fishes[8] = fishes0
		fishes[6] += fishes0
		fmt.Printf("Day %v Fishes %v\n", day+1, fishes)
	}

	totalFishes := 0
	for _, fish := range fishes {
		totalFishes += fish
	}
	fmt.Printf("Total fishes %v\n", totalFishes)
}
