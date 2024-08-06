package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	givenEquasion, error := reader.ReadString('\n')

	if error != nil {
		panic(error)
	}
	trimmedEquasion := strings.TrimSuffix(givenEquasion, "\n")

	firstRune := trimmedEquasion[0] // Should be sufficient

	switch {
	case firstRune == 'I' || firstRune == 'V' || firstRune == 'X':
		fmt.Println(toRoman(calc(parseEqRoman(trimmedEquasion))))
	case '0' <= firstRune && firstRune <= '9':
		result, err := calc(parseEqArabic(trimmedEquasion))
		if err != nil {
			panic(err)
		} else {
			fmt.Println(result)
		}
	default:
		panic(fmt.Errorf("Wrong input: %s", trimmedEquasion))
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

func parseEqRoman(givenEquasion string) (byte, int, int) {
	lexemes := strings.Split(givenEquasion, " ")

	if len(lexemes) != 3 {
		panic(fmt.Errorf("Unable to parse %s", givenEquasion))
	}

	first := parseRomanNum(lexemes[0])
	sign := lexemes[1][0]
	second := parseRomanNum(lexemes[2])

	return sign, first, second
}

func parseRomanNum(givenRomanNum string) int {
	romanMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	romanNum, ok := romanMap[givenRomanNum]

	if !ok {
		panic(fmt.Errorf("Unable to parse Roman number %s", givenRomanNum))
	}

	return romanNum
}

func parseEqArabic(givenEquasion string) (byte, int, int) {
	lexemes := strings.Split(givenEquasion, " ")

	if len(lexemes) != 3 {
		panic(fmt.Errorf("Unable to parse %s", givenEquasion))
	}

	first, err := strconv.Atoi(lexemes[0])
	if err != nil {
		panic(err)
	}
	sign := lexemes[1][0]
	second, err := strconv.Atoi(lexemes[2])
	if err != nil {
		panic(err)
	}

	return sign, first, second
}

func calc(sign byte, first int, second int) (int, error) {

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
