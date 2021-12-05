package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var posX int
	var posY int
	var aim int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var command string
		var steps int
		input := scanner.Text()
		if input == "" {
			break
		}
		fmt.Sscanf(input, "%s %v", &command, &steps)
		switch command {
		case "forward":
			posX += steps
			posY += aim * steps
		case "backward":
			posX -= steps
		case "down":
			aim += steps
		case "up":
			aim -= steps
		}
	}
	fmt.Println(strconv.Itoa(posX * posY))
}
