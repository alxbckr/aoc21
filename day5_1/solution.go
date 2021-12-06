package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type LineSegment struct {
	from, to Point
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []LineSegment{}
	points := make([][]int, 1000)
	for idx := range points {
		points[idx] = make([]int, 1000)
	}

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		var line LineSegment
		fmt.Sscanf(input, "%d,%d -> %d,%d", &line.from.x, &line.from.y, &line.to.x, &line.to.y)
		lines = append(lines, line)
	}

	fmt.Printf("Lines: %v\n", lines)

	for _, line := range lines {
		if line.from.y == line.to.y {
			if line.to.x > line.from.x {
				for x := line.from.x; x <= line.to.x; x++ {
					points[line.from.y][x] += 1
				}
			} else {
				for x := line.to.x; x <= line.from.x; x++ {
					points[line.from.y][x] += 1
				}
			}
		} else if line.from.x == line.to.x {
			if line.to.y > line.from.y {
				for y := line.from.y; y <= line.to.y; y++ {
					points[y][line.from.x] += 1
				}
			} else {
				for y := line.to.y; y <= line.from.y; y++ {
					points[y][line.from.x] += 1
				}
			}
		}
	}

	sum := 0
	for row := range points {
		for col := range points {
			if points[row][col] >= 2 {
				sum++
			}
		}
	}

	fmt.Printf("Points: %v\n", points)

	fmt.Printf("Result: %d\n", sum)
}
