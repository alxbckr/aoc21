package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Stack struct {
	stack *list.List
}

func (c *Stack) Push(value string) {
	c.stack.PushFront(value)
}

func (c *Stack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("pop error: stack is empty")
}

func (c *Stack) Empty() bool {
	return c.stack.Len() == 0
}

func (c *Stack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep error: stack datatype is incorrect")
	}
	return "", fmt.Errorf("peep error: stack is empty")
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}

	codeLines := [][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		line := []string{}
		line = append(line, strings.SplitAfter(input, "")...)
		codeLines = append(codeLines, line)
	}

	lineScores := []int{}
	for _, line := range codeLines {
		lineCorrupted := false
		stack := &Stack{stack: list.New()}
		for _, r := range line {
			if strings.Contains("([{<", r) {
				stack.Push(r)
			} else {
				f, _ := stack.Front()
				exp := ""
				switch f {
				case "(":
					exp = ")"
				case "[":
					exp = "]"
				case "{":
					exp = "}"
				case "<":
					exp = ">"
				}
				if r != exp {
					fmt.Printf("Line corrupted, expected %v, but found %v instead\n", exp, r)
					lineCorrupted = true
					break
				} else {
					stack.Pop()
				}
			}
		}

		if lineCorrupted {
			continue
		}

		completion := ""
		for !stack.Empty() {
			r, _ := stack.Front()
			stack.Pop()
			switch r {
			case "(":
				completion += ")"
			case "[":
				completion += "]"
			case "{":
				completion += "}"
			case "<":
				completion += ">"
			}
		}

		fmt.Printf("Completed line %v\n", completion)

		lineScore := 0
		for _, r := range strings.SplitAfter(completion, "") {
			score := 0
			switch r {
			case ")":
				score = 1
			case "]":
				score = 2
			case "}":
				score = 3
			case ">":
				score = 4
			}
			lineScore = lineScore*5 + score
		}
		fmt.Printf("Total points: %v\n", lineScore)
		lineScores = append(lineScores, lineScore)
	}

	sort.Ints(lineScores)
	fmt.Printf("Result %v\n", lineScores[len(lineScores)/2])
}
