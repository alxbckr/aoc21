package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	const windowSize = 3

	scanner := bufio.NewScanner(os.Stdin)
	windows := make([]int, windowSize+1)
	windowA := 0
	windowB := 1
	index := 1
	result := 0
	for scanner.Scan() {
		if currentMeasure, err := strconv.Atoi(scanner.Text()); err == nil {

			for i := 0; i <= windowSize; i++ {
				if index >= i+1 && (index-i)%(windowSize+1) != 0 {
					windows[i] += currentMeasure
				}
			}

			//fmt.Printf("Windows 1: [%v %v %v %v]\n", windows[0], windows[1], windows[2], windows[3])

			if index > 3 {
				if windows[windowB] > windows[windowA] {
					result++
				}
				//fmt.Printf("Index %v Result %v Window A %v Window B %v\n", index, result, windowA, windowB)
				windowA = (windowA + 1) % (windowSize + 1)
				windowB = (windowB + 1) % (windowSize + 1)
			}

			for i := 0; i <= windowSize; i++ {
				if index >= i+1 && (index-i)%(windowSize+1) == 0 {
					windows[i] = 0
				}
			}

			//fmt.Printf("Windows 2: [%v %v %v %v]\n", windows[0], windows[1], windows[2], windows[3])
			index++
		} else {
			break
		}
	}
	fmt.Println(result)
}
