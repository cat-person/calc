package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	equasion, _ := reader.ReadString('\n')
	fmt.Print(equasion)
}
