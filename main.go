package main

import (
	"fmt"
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

	//d7p1()
	//d7p2()

	//d8p1()
	//d8p2()

	//d9p1()
	//d9p2()

	//d10p1()
	//d10p2()

	d11p1()
	//d11p2()
}

func fileToString(day int) string {
	path := "./inputs/input_day" + dayToString(day) + ".txt"
	rawInput, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(rawInput)
}

func testfileToString(day int) string {
	path := "./inputs/input_day" + dayToString(day) + "_test.txt"
	rawInput, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(rawInput)
}

func dayToString(day int) string {
	if day < 10 {
		return "0" + strconv.Itoa(day)
	}
	return strconv.Itoa(day)
}

func splitLines(s string) []string {
	return strings.Split(s, "\r\n")
}

func splitDoubleLines(s string) []string {
	return strings.Split(s, "\r\n")
}

func printIntMatrix(s [][]int) {
	fmt.Printf("[\n")
	for r := 0; r < len(s); r++ {
		fmt.Printf("[")
		for c := 0; c < len(s[r]); c++ {
			if c != len(s[r])-1 {
				fmt.Printf("%v ", s[r][c])
			}else {
				fmt.Printf("%v", s[r][c])
			}
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("]\n")
}

func printBoolMatrix(s [][]bool) {
	fmt.Printf("[\n")
	for r := 0; r < len(s); r++ {
		fmt.Printf("[")
		for c := 0; c < len(s[r]); c++ {
			if c != len(s[r])-1 {
				if s[r][c] == true {
					fmt.Printf("T ")
				}else{
					fmt.Printf("F ")
				}
			}else {
				if s[r][c] == true {
					fmt.Printf("T")
				}else{
					fmt.Printf("F")
				}
			}
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("]\n")
}