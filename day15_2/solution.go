package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/Workiva/go-datastructures/queue"
)

type Node struct {
	x, y int
}

type Dist struct {
	Node
	dist int
}

func (d Dist) Compare(other queue.Item) int {
	oi := other.(Dist)
	if d.dist == oi.dist {
		return 0
	} else if d.dist > oi.dist {
		return 1
	}
	return -1
}
func getAdjNodes(node Node, maxX, maxY int) []Node {
	res := []Node{}
	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, 1, -1}
	for i := 0; i < 4; i++ {
		n := Node{x: node.x + dx[i], y: node.y + dy[i]}
		if n.x < 0 || n.y < 0 || n.x > maxX || n.y > maxY {
			continue
		}
		res = append(res, n)
	}
	return res
}

func buildMap(weights [][]int, maxX, maxY int) [][]int {
	res := make([][]int, 5*(maxY+1))
	for y := range res {
		res[y] = make([]int, 5*(maxX+1))
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			res[y][x] = weights[y][x]
		}
	}
	// build first row
	for i := 1; i < 5; i++ {
		for y := 0; y <= maxY; y++ {
			for x := (i - 1) * (maxX + 1); x <= (i-1)*(maxX+1)+maxX; x++ {
				newX := x + maxX + 1
				res[y][newX] = res[y][x] + 1
				if res[y][newX] > 9 {
					res[y][newX] = 1
				}
			}
		}
	}
	// build lower rows
	for i := 1; i < 5; i++ {
		for y := (i - 1) * (maxY + 1); y <= (i-1)*(maxY+1)+maxY; y++ {
			for x := range res[y] {
				newY := y + maxY + 1
				res[newY][x] = res[y][x] + 1
				if res[newY][x] > 9 {
					res[newY][x] = 1
				}
			}
		}
	}

	return res
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer f.Close()

	weights := [][]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		line := []int{}
		for _, r := range input {
			line = append(line, int(r-'0'))
		}
		weights = append(weights, line)
	}

	maxX := len(weights[0]) - 1
	maxY := len(weights) - 1

	weights = buildMap(weights, maxX, maxY)
	maxX = len(weights[0]) - 1
	maxY = len(weights) - 1

	dist := make(map[Node]int)
	distQueue := queue.NewPriorityQueue(maxX*maxY, true)
	for y := range weights {
		for x := range weights[y] {
			d := Dist{}
			d.x = x
			d.y = y
			if x != 0 || y != 0 {
				d.dist = math.MaxInt32
			}
			distQueue.Put(d)
		}
	}

	endNode := Node{x: maxX, y: maxY}
	currNode := Node{x: 0, y: 0}

	for !distQueue.Empty() {
		items, _ := distQueue.Get(1)
		currNode = items[0].(Dist).Node
		if currNode == endNode {
			break
		}

		for _, a := range getAdjNodes(currNode, maxX, maxY) {
			alt := dist[currNode] + weights[a.y][a.x]
			if alt < dist[a] || dist[a] == 0 {
				dist[a] = alt
				d := Dist{Node: a, dist: alt}
				distQueue.Put(d)
			}
		}
	}

	fmt.Printf("Path: %v\n", dist[Node{x: maxX, y: maxY}])
}
