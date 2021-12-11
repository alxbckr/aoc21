package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
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

	stack := &Stack{stack: list.New()}
	totalScore := 0
	for _, line := range codeLines {
		for _, r := range line {
			if strings.Contains("([{<", r) {
				stack.Push(r)
			} else {
				f, _ := stack.Front()
				exp := ""
				score := 0
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
				switch r {
				case ")":
					score = 3
				case "]":
					score = 57
				case "}":
					score = 1197
				case ">":
					score = 25137
				}

				if r != exp {
					fmt.Printf("Expected %v, but found %v instead\n", exp, r)
					totalScore += score
					break
				} else {
					stack.Pop()
				}
			}
		}
	}

	fmt.Printf("Score: %v\n", totalScore)
}
