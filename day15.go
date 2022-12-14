package main

import "fmt"

func d15p1() {
	file := fileToString(15)

	lines := splitLines(fileToString(15))
	for l := 0; l < len(lines); l++ {
		fmt.Printf("%v", lines[l])
	}

	fmt.Printf("%v", file)
}

func d15p2() {
	file := fileToString(15)

	lines := splitLines(fileToString(15))
	for l := 0; l < len(lines); l++ {
		fmt.Printf("%v", lines[l])
	}

	fmt.Printf("%v", file)
}