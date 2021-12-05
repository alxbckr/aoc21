package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const numBits int = 12
	numScans := 0
	scanner := bufio.NewScanner(os.Stdin)
	scans := []int{}
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		scan := 0
		fmt.Sscanf(input, "%b", &scan)
		scans = append(scans, scan)
		numScans++
	}

	oxygenScans := scans[:]
	for bitIndex := numBits - 1; bitIndex >= 0; bitIndex-- {
		countBits := 0
		bitMask := 1 << bitIndex
		for _, scan := range oxygenScans {
			if scan&bitMask > 0 {
				countBits++
			}
		}

		oxygenScansTmp := []int{}
		for _, scan := range oxygenScans {
			if countBits >= len(oxygenScans)-countBits && scan&bitMask > 0 ||
				countBits < len(oxygenScans)-countBits && scan&bitMask == 0 {
				oxygenScansTmp = append(oxygenScansTmp, scan)
			}
		}
		oxygenScans = oxygenScansTmp[:]
		if len(oxygenScans) == 1 {
			break
		}
	}

	co2ScrubberScans := scans[:]
	for bitIndex := numBits - 1; bitIndex >= 0; bitIndex-- {
		countBits := 0
		bitMask := 1 << bitIndex
		for _, scan := range co2ScrubberScans {
			if scan&bitMask > 0 {
				countBits++
			}
		}
		co2scrubberScansTmp := []int{}
		for _, scan := range co2ScrubberScans {
			if countBits < len(co2ScrubberScans)-countBits && scan&bitMask > 0 ||
				countBits >= len(co2ScrubberScans)-countBits && scan&bitMask == 0 {
				co2scrubberScansTmp = append(co2scrubberScansTmp, scan)
			}
		}
		co2ScrubberScans = co2scrubberScansTmp[:]
		if len(co2ScrubberScans) == 1 {
			break
		}
	}

	fmt.Printf("Oxygen generator rating: %v\n", oxygenScans[0])
	fmt.Printf("CO2 scrubber rating: %v\n", co2ScrubberScans[0])
	fmt.Printf("life support rating: %v\n", co2ScrubberScans[0]*oxygenScans[0])
}
