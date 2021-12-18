package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct {
	x, y int
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
	dist := make(map[Node]int)
	visited := make(map[Node]bool)

	minNode := Node{x: 0, y: 0}
	currNode := Node{x: 0, y: 0}

	for len(visited) < (maxX+1)*(maxY+1) {
		visited[currNode] = true
		for _, a := range getAdjNodes(currNode, maxX, maxY) {
			alt := dist[currNode] + weights[a.y][a.x]
			if alt < dist[a] || dist[a] == 0 {
				dist[a] = alt
			}
		}

		minVal := math.MaxInt32
		for k, v := range dist {
			if v < minVal && v != 0 && !visited[k] {
				minNode = k
				minVal = v
			}
		}
		currNode = minNode
	}

	fmt.Printf("Path: %v\n", dist[Node{x: maxX, y: maxY}])
}
