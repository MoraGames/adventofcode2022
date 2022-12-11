package main

import (
	"fmt"
	"strconv"
	"strings"
)

func d10p1() {
	var currentCycle int //not completed yet
	var sumSignalStrength, xRegister int = 0, 1

	lines := splitLines(fileToString(10))
	for l := 0; l < len(lines); l++ {
		line := strings.Split(lines[l], " ")
		fmt.Printf("\n\n[DEBUG] ------------------------)")
		fmt.Printf("\n\tline = %v\n\tcurrentCycle = %v\n\txRegister = %v\n\tsumSignalStrength = %v", line, currentCycle, xRegister, sumSignalStrength)

		if line[0] == "addx" {
			currentCycle++
			readSignal(&sumSignalStrength, currentCycle, xRegister)

			currentCycle++
			readSignal(&sumSignalStrength, currentCycle, xRegister)

			value, err := strconv.Atoi(line[1])
			if err != nil {
				panic(err)
			}
			xRegister += value
		} else if line[0] == "noop" { //noop
			currentCycle++
			readSignal(&sumSignalStrength, currentCycle, xRegister)
		}
	}

	fmt.Printf("\n\n")
	fmt.Printf("sumSignalStrength = %d\n", sumSignalStrength)
}

func d10p2() {
	var currentCycle int //not completed yet
	var sumSignalStrength, xRegister int = 0, 1
	var CTR string

	lines := splitLines(fileToString(10))
	for l := 0; l < len(lines); l++ {
		line := strings.Split(lines[l], " ")
		fmt.Printf("\n\n[DEBUG] ------------------------)")
		fmt.Printf("\n\tline = %v\n\tcurrentCycle = %v\n\txRegister = %v\n\tsumSignalStrength = %v", line, currentCycle, xRegister, sumSignalStrength)

		if line[0] == "addx" {
			writePixel(&CTR, currentCycle, xRegister)
			currentCycle++
			readSignal(&sumSignalStrength, currentCycle, xRegister)

			writePixel(&CTR, currentCycle, xRegister)
			currentCycle++
			readSignal(&sumSignalStrength, currentCycle, xRegister)

			value, err := strconv.Atoi(line[1])
			if err != nil {
				panic(err)
			}
			xRegister += value
		} else if line[0] == "noop" { //noop
			writePixel(&CTR, currentCycle, xRegister)
			currentCycle++
			readSignal(&sumSignalStrength, currentCycle, xRegister)
		}
	}

	fmt.Printf("\n\n")
	for i := 0; i < 240; i++ {
		if i%40 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", CTR[i])
	}
}

func readSignal(sumSignalStrength *int, currentCycle, xRegister int) {
	switch currentCycle {
	case 20:
		fallthrough
	case 60:
		fallthrough
	case 100:
		fallthrough
	case 140:
		fallthrough
	case 180:
		fallthrough
	case 220:
		fmt.Printf("\n\t%vth cycle!\n\t(sumSignalStrength = %v | currentCycle = %v | xRegister %v)", currentCycle, *sumSignalStrength, currentCycle, xRegister)
		*sumSignalStrength += currentCycle * xRegister
	}
}

func writePixel(CTR *string, currentCycle, xRegister int) {
	if currentCycle%40 == xRegister-1 || currentCycle%40 == xRegister || currentCycle%40 == xRegister+1 {
		*CTR += "#"
	} else {
		*CTR += "."
	}
}
