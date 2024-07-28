package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type calculateFunc = func([][]byte) (int, bool)
type formatFunc = func(int, bool) string

func getInput(entryText string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(entryText)
	scanner.Scan()

	return strings.TrimSpace(strings.ToLower(scanner.Text()))
}

func parseInput(text []byte) [][]byte {
	re := regexp.MustCompile(`(now)|\s(\+|-)\s|[0-9.:/-]*`)
	arr := re.FindAll(text, -1)

	return arr
}

func main() {
	fmt.Println("Datetime Calculator")

	var run bool = true
	var menu bool = true

	var calculate calculateFunc
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
			result, formatType := calculate(exp)
			fmt.Printf("|- %s\n", format(result, formatType))
		}
	}
}
