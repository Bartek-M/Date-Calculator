package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("time > ")
	scanner.Scan()

	return scanner.Text()
}

func parseInput(text []byte) [][]byte {
	// datePattern := `([0-2][\d]|3[0-1])(/|-|.)(0[1-9]|1[0-2])(/|-|.)(\d{4})`
	re := regexp.MustCompile(`(\+|-)|[0-9:./]*`)
	arr := re.FindAll(text, -1)

	return arr
}

func main() {
	fmt.Print("Datetime Calculator\n\n")

	text := getInput()
	exp := parseInput([]byte(text))

	result := calculate(exp)
	fmt.Printf("|- %s\n\n", formatTime(result))
}
