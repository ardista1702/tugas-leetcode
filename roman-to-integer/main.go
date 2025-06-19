package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4,
		1,
	}
	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I",
	}

	result := []string{}

	for i := range values {
		for num >= values[i] {
			num -= values[i]
			result = append(result, symbols[i])
		}
	}

	return strings.Join(result,"")
}

func main() {
	int_ := 3749
	res := intToRoman(int_)
	fmt.Println("res", res)
}