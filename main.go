package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	//d1p1()
	//d1p2()

	//d2p1()
	//d2p2()

	//d3p1()
	//d3p2()

	//d4p1()
	//d4p2()

	//d5p1()
	//d5p2()

	//d6p1()
	//d6p2()

	d7p1()
	//d7p2()
}

func fileToString(day int) string {
	path := "./inputs/input_day" + strconv.Itoa(day) + ".txt"
	rawInput, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(rawInput)
}

func splitLines(s string) []string {
	return strings.Split(s, "\r\n")
}

func splitDoubleLines(s string) []string {
	return strings.Split(s, "\r\n")
}