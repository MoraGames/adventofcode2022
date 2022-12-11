package main

import "fmt"

func d6p1() {
	file := fileToString(6)
	for i := 0; i < len(file); i++ {
		section := file[i : i+4]
		//fmt.Printf("Section: %v\n", section)

		var exit bool
		for s1 := 0; s1 < 4 && exit == false; s1++ {
			for s2 := s1 + 1; s2 < 4 && exit == false; s2++ {
				if s1 != s2 && section[s1] == section[s2] {
					exit = true
				}
			}
		}
		if exit == false {
			fmt.Printf("Section[%d:%d] = %v\n", i, i+4, section)
			break
		}
	}
}

func d6p2() {
	file := fileToString(6)
	for i := 0; i < len(file); i++ {
		section := file[i : i+14]
		//fmt.Printf("Section: %v\n", section)

		var exit bool
		for s1 := 0; s1 < 14 && exit == false; s1++ {
			for s2 := s1 + 1; s2 < 14 && exit == false; s2++ {
				if s1 != s2 && section[s1] == section[s2] {
					exit = true
				}
			}
		}
		if exit == false {
			fmt.Printf("Section[%d:%d] = %v\n", i, i+14, section)
			break
		}
	}
}
