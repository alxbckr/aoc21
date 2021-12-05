package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func printBoard(board [][]int) {
	for _, row := range board {
		for _, elem := range row {
			fmt.Printf("%3v", elem)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func updateBoardMarkers(boards [][][]int, boardMarkers [][][]int, randomNumber int) {
	for idxBoard, board := range boards {
		for idxRow, row := range board {
			for idxNum, num := range row {
				if num == randomNumber {
					boardMarkers[idxBoard][idxRow][idxNum] = 1
				}
			}
		}
	}
}

func getWinningBoardIdx(boardMarkers [][][]int, winningBoards []int) int {
	for idxBoard, board := range boardMarkers {
		if winningBoards[idxBoard] > 0 {
			continue
		}
		for _, row := range board {
			summRow := 0
			for _, num := range row {
				summRow += num
			}
			if summRow == 5 {
				return idxBoard
			}
		}
		for col := 0; col < 5; col++ {
			summCol := 0
			for _, row := range board {
				summCol += row[col]
			}
			if summCol == 5 {
				return idxBoard
			}
		}
	}
	return -1
}

func getWinningBoardSum(boards [][][]int, boardMarkers [][][]int, boardIdx int) int {
	summ := 0
	for idxRow, row := range boards[boardIdx] {
		for idxCol, num := range row {
			if boardMarkers[boardIdx][idxRow][idxCol] == 0 {
				summ += num
			}
		}
	}
	return summ
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	randomInput := []int{}
	boards := [][][]int{}
	index := 0

	boardIndex := 0
	for scanner.Scan() {
		input := scanner.Text()
		if input == "-" {
			break
		}
		if index == 0 {
			for _, str := range strings.Split(input, ",") {
				if num, err := strconv.Atoi(str); err == nil {
					randomInput = append(randomInput, num)
				}
			}
		} else {
			if input == "" {
				if index > 2 {
					boardIndex++
				}
				board := make([][]int, 0)
				boards = append(boards, board)
				continue
			}
			boardRow := make([]int, 0)
			inputPrepared := strings.TrimSpace(input)
			inputPrepared = regexp.MustCompile(`\s+`).ReplaceAllString(inputPrepared, " ")
			for _, str := range strings.Fields(inputPrepared) {
				if num, err := strconv.Atoi(str); err == nil {
					boardRow = append(boardRow, num)
				}
			}
			boards[boardIndex] = append(boards[boardIndex], boardRow)
		}
		index++
	}

	boardMarkers := make([][][]int, len(boards))
	for idxB := range boardMarkers {
		boardMarkers[idxB] = make([][]int, 5)
		for idxR := range boardMarkers[idxB] {
			boardMarkers[idxB][idxR] = make([]int, 5)
		}
	}

	lastWinBoardIdx := -1
	lastWinNumber := -1
	lastWinBoardSum := 0
	winnerBoards := make([]int, len(boards))
	numWinnerBoards := 0
	for _, randomNuber := range randomInput {
		updateBoardMarkers(boards, boardMarkers, randomNuber)
		for {
			boardIdx := getWinningBoardIdx(boardMarkers, winnerBoards)
			if boardIdx < 0 {
				break
			}
			lastWinBoardIdx = boardIdx
			lastWinNumber = randomNuber
			lastWinBoardSum = getWinningBoardSum(boards, boardMarkers, lastWinBoardIdx)
			winnerBoards[boardIdx] = 1
			fmt.Printf("Board %v wins\n", boardIdx)
			fmt.Printf("Number %v\n", lastWinNumber)
			numWinnerBoards++

			printBoard(boards[lastWinBoardIdx])
			printBoard(boardMarkers[lastWinBoardIdx])
		}
		if numWinnerBoards == len(boards) {
			break
		}
	}
	fmt.Printf("BoardIdx: %v\n", lastWinBoardIdx)
	fmt.Printf("Number: %v\n", lastWinNumber)
	fmt.Printf("Result: %v\n", lastWinBoardSum*lastWinNumber)
}
