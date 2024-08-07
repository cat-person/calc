package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	CheckEquasion("1 + 2", "3")
	CheckEquasion("VI / III", "II")
	CheckEquasion("10 * 10", "100")
	CheckEquasion("X * X", "C")
	CheckEquasion("X / III", "III")
	CheckEquasion("X / X", "I")
	CheckEquasion("1 / 10", "0")

	CheckPanic("I - II")
	CheckPanic("I + 1")
	CheckPanic("3 - II")
	CheckPanic("IV * 5")
	CheckPanic("I / IX")
	CheckPanic("1 + 2 + 3")
	CheckPanic("-1 + -2")
}

func CheckEquasion(givenEquasion string, expectedResult string) {
	actualResult, err := SolveEq(givenEquasion)
	if err != nil {
		fmt.Printf(">>> TEST ERROR: attempting solving equasion: %s caused error: %e \n", givenEquasion, err)
	} else if actualResult != expectedResult {
		fmt.Printf(">>> TEST ERROR: calculation error while solving equasion: %s: expectedResult: %s actualResult: %s \n", givenEquasion, expectedResult, actualResult)
	} else {
		fmt.Printf("PASSED: equasion: %s: is solved with expectedResult: %s \n", givenEquasion, expectedResult)
	}
}

func CheckPanic(givenEquasion string) {
	result, err := SolveEq(givenEquasion)

	if err == nil {
		fmt.Printf(">>> TEST ERROR: equasion: %s has been solved with the result %s but error was expected\n", givenEquasion, result)
	} else {
		fmt.Printf("PASSED: equasion: %s: returned error: %s. Just as planned\n", givenEquasion, err)
	}
}
