package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type calcFunc = func(exp [][]byte) int
type formatFunc = func(int) string

func getInput(entryText string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(entryText)
	scanner.Scan()

	return strings.TrimSpace(strings.ToLower(scanner.Text()))
}

func parseInput(text []byte) [][]byte {
	// datePattern := `([0-2][\d]|3[0-1])(/|-|.)(0[1-9]|1[0-2])(/|-|.)(\d{4})`
	re := regexp.MustCompile(`(\+|-)|[0-9:./]*`)
	arr := re.FindAll(text, -1)

	return arr
}

func main() {
	fmt.Print("Datetime Calculator\n")

	var run = true
	var menu bool = true

	var calculate calcFunc
	var format formatFunc
	var entry string = "\n[time | date] > "

	for run {
		inpt := getInput(entry)

		if menu {
			if inpt != "time" && inpt != "date" {
				continue
			}

			inpt = fmt.Sprintf("/%s", inpt)
			menu = false
		}

		switch inpt {
		case "/time":
			calculate = calculateTime
			format = formatTime
			entry = "\ntime > "
		case "/date":
			calculate = calculateDate
			format = formatDate
			entry = "\ndate > "
		case "quit", "exit":
			run = false
		default:
			exp := parseInput([]byte(inpt))
			result := calculate(exp)
			fmt.Printf("|- %s\n", format(result))
		}
	}
}
