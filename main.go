package main

import "fmt"

func checkRevenge(s string) string {
	shots := 0
	revenge := 0

	for _, char := range s {
		if char == 'S' {
			shots++
		} else if char == 'R' {
			if shots == 0 {
				return "Bad boy"
			}
			revenge++
			shots--
		}
	}

	if revenge >= shots {
		return "Good boy"
	} else {
		return "Bad boy"
	}
}

func main() {
	input := "SRSSRR"
	result := checkRevenge(input)
	fmt.Println(result)

	input2 := "RSSRR"
	result2 := checkRevenge(input2)
	fmt.Println(result2)

	input3 := "SSSRRRRS"
	result3 := checkRevenge(input3)
	fmt.Println(result3)

	input4 := "SSRR"
	result4 := checkRevenge(input4)
	fmt.Println(result4)

	input5 := "SRRSSR"
	result5 := checkRevenge(input5)
	fmt.Println(result5)
}
