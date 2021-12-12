package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func flash(x, y int, octopuses [][]int, flashedMap map[coord]bool) {
	flashedMap[coord{x: x, y: y}] = true
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			yA := y + dy
			xA := x + dx
			if yA == y && xA == x ||
				xA < 0 || yA < 0 || xA == len(octopuses) || yA == len(octopuses) {
				continue
			}
			octopuses[yA][xA] += 1
		}
	}
}

func checkFlashed(octopuses [][]int, flashedMap map[coord]bool) bool {
	flashed := false
	for y := range octopuses {
		for x := range octopuses[y] {
			if octopuses[y][x] > 9 && !flashedMap[coord{x: x, y: y}] {
				flash(x, y, octopuses, flashedMap)
				flashed = true
			}
		}
	}

	return flashed
}

func resetFlashed(octopuses [][]int) {
	for y := range octopuses {
		for x := range octopuses[y] {
			if octopuses[y][x] > 9 {
				octopuses[y][x] = 0
			}
		}
	}
}

func synced(octopuses [][]int) bool {
	for y := range octopuses {
		for x := range octopuses[y] {
			if octopuses[y][x] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}

	octopuses := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		octopusesLine := []int{}
		for _, r := range strings.SplitAfter(input, "") {
			num, _ := strconv.Atoi(r)
			octopusesLine = append(octopusesLine, num)
		}
		octopuses = append(octopuses, octopusesLine)
	}

	step := 0
	for step = 0; step < 1000; step++ {
		flashedMap := make(map[coord]bool)

		for y := range octopuses {
			for x := range octopuses[y] {
				octopuses[y][x] += 1
			}
		}

		for checkFlashed(octopuses, flashedMap) {
		}
		resetFlashed(octopuses)
		if synced(octopuses) {
			break
		}
	}

	fmt.Printf("Sync step %v\n", step+1)
}
