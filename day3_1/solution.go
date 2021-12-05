package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bitcount := make([]int, 12)
	counter := 0
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		counter++
		number := 0
		fmt.Sscanf(input, "%b", &number)
		for i := 0; i < 12; i++ {
			if number&(1<<i) > 0 {
				bitcount[i] += 1
			}
		}
	}
	fmt.Printf("%v\n", bitcount)
	gammaRate := 0
	for i := 0; i < 12; i++ {
		if bitcount[i] > counter-bitcount[i] {
			gammaRate |= 1 << i
		}
	}
	fmt.Printf("Gamma: %b\n", gammaRate)
	epsilonRate := 0xfff & ^gammaRate
	fmt.Printf("Epsilon %b\n", epsilonRate)
	fmt.Println(gammaRate * epsilonRate)
}
