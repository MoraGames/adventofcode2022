package main

import "fmt"

func d12p1() {
	file := fileToString(12)

	lines := splitLines(fileToString(12))
	for l := 0; l < len(lines); l++ {
		fmt.Print(lines[l])
	}

	fmt.Print(file, lines)
}

func d12p2() {
	file := fileToString(12)

	lines := splitLines(fileToString(12))
	for l := 0; l < len(lines); l++ {
		fmt.Print(lines[l])
	}

	fmt.Print(file, lines)
}