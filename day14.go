package main

import "fmt"

func d14p1() {
	file := fileToString(14)

	lines := splitLines(fileToString(14))
	for l := 0; l < len(lines); l++ {
		fmt.Print(lines[l])
	}

	fmt.Print(file, lines)
}

func d14p2() {
	file := fileToString(14)

	lines := splitLines(fileToString(14))
	for l := 0; l < len(lines); l++ {
		fmt.Print(lines[l])
	}

	fmt.Print(file, lines)
}
