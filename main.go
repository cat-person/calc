package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')

	if error != nil {
		panic(error)
	}

	switch input[0] {
	case 'I', 'V':
		fmt.Println(toRoman(calc(parseRoman(input))))
	}
	panic(fmt.Errorf("Wrong input: %s", input))
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
		return "V"
	}
}

func parseRoman(string) (rune, int, int) {
	return '+', 2, 2
}

func parseArabic(string) (rune, int, int) {
	return '+', 2, 2
}

func calc(sign rune, first int, second int) (int, error) {
	switch sign {
	case '+':
		return -1, nil
	case '-':
		return -1, nil
	case '*':
		return -1, nil
	case '/':
		return -1, nil
	default:
		return -1, fmt.Errorf("Unable to recognise sign %s", sign)
	}
}
