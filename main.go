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
	userInput, error := reader.ReadString('\n')

	if error != nil {
		panic(error)
	}
	trimmedUserInput := strings.TrimSuffix(userInput, "\n")

	result, error := SolveEq(trimmedUserInput)

	if error != nil {
		panic(error)
	}

	fmt.Println(result)
}

func SolveEq(givenEquasion string) (string, error) {

	firstRune := givenEquasion[0] // Should be sufficient

	switch {
	case firstRune == 'I' || firstRune == 'V' || firstRune == 'X':
		return toRoman(calc(SolveEqRoman(givenEquasion)))
	case '0' <= firstRune && firstRune <= '9':
		result, err := calc(SolveEqArabic(givenEquasion))
		if err == nil {
			return strconv.Itoa(result), nil
		} else {
			return "", err
		}
	default:
		return "", fmt.Errorf("wrong input: %s", givenEquasion)
	}
}

func toRoman(givenResult int, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if givenResult <= 0 {
		if givenResult == 0 {
			return "", fmt.Errorf("result is zero")
		} else {
			return "", fmt.Errorf("result is negative")
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

		return result.String(), nil
	}
}

func SolveEqRoman(givenEquasion string) (byte, int, int, error) {
	lexemes := strings.Split(givenEquasion, " ")

	if len(lexemes) != 3 {
		return '@', 0, 0, fmt.Errorf("unable to parse %s", givenEquasion)
	}

	first, err := ParseRomanNum(lexemes[0])
	if err != nil {
		return '@', 0, 0, err
	}
	sign := lexemes[1][0]
	second, err := ParseRomanNum(lexemes[2])
	if err != nil {
		return '@', 0, 0, err
	}

	return sign, first, second, nil
}

func ParseRomanNum(givenRomanNum string) (int, error) {
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
		return 0, fmt.Errorf("unable to parse %s as roman number", givenRomanNum)
	}
	return romanNum, nil
}

func SolveEqArabic(givenEquasion string) (byte, int, int, error) {
	lexemes := strings.Split(givenEquasion, " ")

	if len(lexemes) != 3 {
		return '@', 0, 0, fmt.Errorf("unable to parse")
	}

	first, err := strconv.Atoi(lexemes[0])
	if err != nil {
		return '@', 0, 0, err
	}
	sign := lexemes[1][0]
	second, err := strconv.Atoi(lexemes[2])
	if err != nil {
		return '@', 0, 0, err
	}

	return sign, first, second, nil
}

func calc(sign byte, first int, second int, err error) (int, error) {
	if err != nil {
		return -1, err
	}

	if first < 0 || first > 10 {
		return -1, fmt.Errorf("first argument: %d is outside of allowed borders [0, 10]", first)
	}

	if second < 0 || second > 10 {
		return -1, fmt.Errorf("second argument: %d is outside of allowed borders [0, 10]", second)
	}

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
		return -1, fmt.Errorf("unable to recognise sign %c", sign)
	}
}
