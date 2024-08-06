package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')

	if error != nil {
		panic(error)
	}

	firstRune := input[0] // Should be sufficient

	switch {
	case firstRune == 'I', firstRune == 'V':
		fmt.Println(toRoman(calc(parseRoman(input))))
	case '0' <= firstRune && firstRune <= '9':
		result, err := calc(parseArabic(input))
		if err != nil {
			panic(err)
		} else {
			fmt.Println(result)
		}
	default:
		panic(fmt.Errorf("Wrong input: %s", input))
	}
}

func parse() (string, error) {
	return "a", nil
}

func toRoman(givenResult int, err error) string {
	if err != nil {
		panic(err)
	}
	if givenResult <= 0 {
		if givenResult == 0 {
			panic("Result is zero")
		} else {
			panic("Result is negative")
		}
	} else {
		vals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
		numerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

		var result strings.Builder

		for i := 0; i < len(vals); i++ {
			for givenResult >= vals[i] {
				givenResult -= vals[i]
				result.WriteString(numerals[i])
			}
		}

		return result.String()
	}
}

func parseRoman(string) (rune, int, int) {

	return '+', 2, 2
}

func parseArabic(givenEquasion string) (rune, int, int) {
	first := 0
	second := 0
	sign := '@'

	magicIdx := 0

	for _, char := range givenEquasion {
		switch {
		case '0' <= char && char <= '9':
			if magicIdx == 0 {
				first = 10*first + int(char) - 48
			} else if magicIdx == 2 {
				second = 10*second + int(char) - 48
			} else {
				panic(fmt.Errorf("Unable to parse equasion %s", givenEquasion))
			}
		case char == '+' || char == '-' || char == '*' || char == '/':
			if magicIdx == 1 {
				sign = char
			} else {
				panic(fmt.Errorf("Unable to parse equasion %s", givenEquasion))
			}
		case char == ' ':
			magicIdx++ // string shouldnt end here
		}
	}

	if magicIdx != 2 {
		panic(fmt.Errorf("Unable to parse equasion %s", givenEquasion))
	}

	return sign, first, second
}

func calc(sign rune, first int, second int) (int, error) {

	switch sign {
	case '+':
		return first + second, nil
	case '-':
		return first - second, nil
	case '*':
		return first * second, nil
	case '/':
		return first / second, nil
	default:
		return -1, fmt.Errorf(`Unable to recognise sign %c`, sign)
	}
}
