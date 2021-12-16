package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputCoords int = iota
	inputInstructions
)

const (
	coordX string = "x"
	coordY string = "y"
)

type FoldInstruction struct {
	coord string
	value int
}

type DotCoord struct {
	x, y int
}

type Matrix [][]int

func NewMatrix(x, y int) Matrix {
	matrix := Matrix{}
	for i := 0; i < y; i++ {
		row := make([]int, x)
		matrix = append(matrix, row)
	}
	return matrix
}

func fold(matrix Matrix, instr FoldInstruction) Matrix {
	newSizeY := len(matrix)
	newSizeX := len(matrix[0])
	if instr.coord == coordX {
		// vertical fold
		newSizeX = instr.value
	} else {
		newSizeY = instr.value
	}
	matrixCopy := NewMatrix(newSizeX, newSizeY)

	if instr.coord == coordX {
		colTo := 0
		for colFrom := len(matrix[0]) - 1; colFrom > instr.value; colFrom-- {
			for y := 0; y < len(matrix); y++ {
				if matrix[y][colTo] != 0 || matrix[y][colFrom] != 0 {
					matrixCopy[y][colTo] = 1
				}
			}
			colTo++
		}
	} else {
		rowTo := 0
		for rowFrom := len(matrix) - 1; rowFrom > instr.value; rowFrom-- {
			for x := 0; x < len(matrix[rowFrom]); x++ {
				if matrix[rowTo][x] != 0 || matrix[rowFrom][x] != 0 {
					matrixCopy[rowTo][x] = 1
				}
			}
			rowTo++
		}
	}

	return matrixCopy
}

func printMatrix(matrix Matrix) {
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	mode := inputCoords
	maxCoord := DotCoord{x: 0, y: 0}
	dotCoords := []DotCoord{}
	foldInstructions := []FoldInstruction{}
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			mode = inputInstructions
			continue
		}

		if mode == inputCoords {
			strs := strings.Split(input, ",")
			x, _ := strconv.Atoi(strs[0])
			y, _ := strconv.Atoi(strs[1])
			dotCoords = append(dotCoords, DotCoord{x: x, y: y})
			if x > maxCoord.x {
				maxCoord.x = x
			}
			if y > maxCoord.y {
				maxCoord.y = y
			}
		} else {
			instr := FoldInstruction{coord: "", value: 0}
			input = strings.Replace(strings.Trim(input, "fold along"), "=", " ", 1)
			fmt.Sscanf(input, "%s %d", &instr.coord, &instr.value)
			foldInstructions = append(foldInstructions, instr)
		}
	}

	matrix := NewMatrix(maxCoord.x+1, maxCoord.y+1)
	for _, coord := range dotCoords {
		matrix[coord.y][coord.x] = 1
	}
	//printMatrix(matrix)

	for i := 0; i < len(foldInstructions); i++ {
		matrix = fold(matrix, foldInstructions[i])
	}
	printMatrix(matrix)

	countDots := 0
	for y := range matrix {
		for x := range matrix[0] {
			countDots += matrix[y][x]
		}
	}
	fmt.Printf("Count dots %v\n", countDots)
}
