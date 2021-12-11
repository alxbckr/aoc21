package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func isLocalMin(matrix [][]int, i, j int) bool {
	matrixWidth := len(matrix[0])
	matrixHeight := len(matrix)
	if (j < matrixWidth-1 && matrix[i][j] < matrix[i][j+1] || j == matrixWidth-1) &&
		(j > 0 && matrix[i][j] < matrix[i][j-1] || j == 0) &&
		(i < matrixHeight-1 && matrix[i][j] < matrix[i+1][j] || i == matrixHeight-1) &&
		(i > 0 && matrix[i][j] < matrix[i-1][j] || i == 0) {

		return true
	}
	return false
}

func countBasinPoints(matrix [][]int, i, j, prevPoint int, v map[Point]bool) int {
	if i < 0 || j < 0 || i == len(matrix) || j == len(matrix[0]) ||
		matrix[i][j] == 9 || matrix[i][j] <= prevPoint || v[Point{x: j, y: i}] {
		return 0
	}
	v[Point{x: j, y: i}] = true
	return 1 + countBasinPoints(matrix, i+1, j, matrix[i][j], v) +
		countBasinPoints(matrix, i-1, j, matrix[i][j], v) +
		countBasinPoints(matrix, i, j+1, matrix[i][j], v) +
		countBasinPoints(matrix, i, j-1, matrix[i][j], v)
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	matrix := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		matrixRow := make([]int, len(input))
		for idx, chr := range strings.SplitAfter(input, "") {
			matrixRow[idx], _ = strconv.Atoi(chr)
		}
		matrix = append(matrix, matrixRow)
	}

	riskLevels := []int{}
	basinPoints := []int{}
	visitedPoints := make(map[Point]bool)
	for i := range matrix {
		for j := range matrix[i] {
			if isLocalMin(matrix, i, j) {
				riskLevels = append(riskLevels, matrix[i][j]+1)
				fmt.Printf("Risk level %v, points %v, %v\n", matrix[i][j], i, j)
				basinPointsCnt := countBasinPoints(matrix, i, j, -1, visitedPoints)
				basinPoints = append(basinPoints, basinPointsCnt)
			}
		}
	}

	sum := 0
	for _, risk := range riskLevels {
		sum += risk
	}

	sort.Slice(basinPoints, func(i, j int) bool {
		return basinPoints[i] > basinPoints[j]
	})

	fmt.Printf("Risk levels %v, Sum %v\n", riskLevels, sum)
	fmt.Printf("Basin points %v\n", basinPoints)
	fmt.Printf("Answer %v\n", basinPoints[0]*basinPoints[1]*basinPoints[2])

}
