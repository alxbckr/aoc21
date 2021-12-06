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
			dirX := 1
			if line.to.x < line.from.x {
				dirX = -1
			}
			x := line.from.x
			for {
				points[line.from.y][x] += 1
				if x == line.to.x {
					break
				}
				x += dirX
			}
		} else if line.from.x == line.to.x {
			dirY := 1
			if line.to.y < line.from.y {
				dirY = -1
			}
			y := line.from.y
			for {
				points[y][line.from.x] += 1
				if y == line.to.y {
					break
				}
				y += dirY
			}
		} else {
			dirX := 1
			if line.to.x < line.from.x {
				dirX = -1
			}
			dirY := 1
			if line.to.y < line.from.y {
				dirY = -1
			}
			y := line.from.y
			x := line.from.x
			for {
				points[y][x] += 1
				if x == line.to.x {
					break
				}
				x += dirX
				y += dirY
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
