package main

import "fmt"

func main() {
	in := "III"
	fmt.Println(romanToInt(in))
}

func romanToInt(s string) int {
	var value = map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	var sum int
	var last rune

	for i, r := range s {
		v, l := value[r], value[last]
		sum += v
		last = r
		if i > 0 && v > l {
			sum -= l * 2
		}
	}

	return sum
}
