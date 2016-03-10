package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func GenerateRoman(input int, alphabet map[int]string, lst []int) string {
	var result string
	for num := range lst {
		num = lst[num]
		count := input / num
		for i := 0; i < count; i++ {
			result = strings.Join([]string{result, alphabet[num]}, "")
		}
		input -= num * count
	}
	return result
}

func main() {
	roman_alphabet := map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD",
		100: "C", 90: "XC", 50: "L", 40: "XL",
		10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}

	var keys []int

	for k := range roman_alphabet {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var input int
	fmt.Println("Hello,\ngive a number grater then 0\nIf input is equal to 0 then \nthe application will be stopped\n")
	for true {
		fmt.Print("Number: ")
		fmt.Scanln(&input)
		if input != 0 {
			fmt.Println(GenerateRoman(input, roman_alphabet, keys))
		} else {
			os.Exit(1)
		}
	}
}
