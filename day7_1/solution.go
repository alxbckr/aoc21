package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	crabs := make(map[int]int)
	crabMax := 0
	for _, str := range strings.Split(scanner.Text(), ",") {
		crabPos, _ := strconv.Atoi(str)
		crabs[crabPos] += 1
		if crabPos > crabMax {
			crabMax = crabPos
		}
	}

	pos := 0
	minFuel := 0
	minPos := 0
	for pos = 0; pos <= crabMax; pos++ {
		fuel := 0
		for crabPos, crabs := range crabs {
			fuel += crabs * Abs(crabPos-pos)
		}
		if fuel < minFuel || minFuel == 0 {
			minFuel = fuel
			minPos = pos
		}
		fmt.Printf("Distance %v Fuel %v\n", pos, fuel)
	}
	fmt.Printf("Distance %v Fuel %v\n", minPos, minFuel)
}
