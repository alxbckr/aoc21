package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	for i := range matrix {
		for j := range matrix[i] {
			if isLocalMin(matrix, i, j) {
				riskLevels = append(riskLevels, matrix[i][j]+1)
				fmt.Printf("Risk level %v, points %v, %v\n", matrix[i][j], i, j)
			}
		}
	}

	sum := 0
	for _, risk := range riskLevels {
		sum += risk
	}

	fmt.Printf("Risk levels %v, Sum %v\n", riskLevels, sum)

}
