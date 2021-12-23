package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type Area struct {
	tl, br Point
}

func calcXSpeedRange(area Area) (int, int) {
	minX := 0
	maxX := 0
	x := 0
	for {
		x += 1
		res := 0
		for i := x; i >= 0; i-- {
			res += i
			if res >= area.tl.x && res <= area.br.x {
				if minX == 0 {
					minX = x
				}
				maxX = x
			}
		}
		if x > area.br.x {
			return minX, maxX
		}
	}
}

func simulate(x, y int, area Area) int {
	dx := x
	dy := y
	px := 0
	py := 0
	topY := 0
	for {
		px += dx
		py += dy
		if py > topY {
			topY = py
		}
		if px >= area.tl.x && px <= area.br.x &&
			py >= area.br.y && py <= area.tl.y {
			fmt.Printf("Speed found x=%d y=%d\n", x, y)
			return topY
		} else if px > area.br.x || py < area.br.y {
			return 0
		}
		if dx > 0 {
			dx -= 1
		}
		dy -= 1
	}
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	area := Area{}
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &area.tl.x, &area.br.x, &area.br.y, &area.tl.y)
	}

	minX, maxX := calcXSpeedRange(area)
	fmt.Printf("minX = %d maxX = %d\n", minX, maxX)

	topY := 0
	topSpeed := Point{}
	for x := minX; x <= maxX; x++ {
		for y := -300; y <= 300; y++ {
			ym := simulate(x, y, area)
			if ym > topY {
				topY = ym
				topSpeed.x = x
				topSpeed.y = y
			}
		}
	}

	fmt.Printf("Highest point %d shoot at %v", topY, topSpeed)

}
