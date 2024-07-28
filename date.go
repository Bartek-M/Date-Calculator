package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseDate(val []byte) (int, int, int) {
	re := regexp.MustCompile(`^(([0-2][\d]|3[0-1])(/|-|.)(0[1-9]|1[0-2])(/|-|.)(\d{1,5}))|\d+$`)
	match := re.FindAll(val, -1)

	if len(match) != 1 {
		return 0, 0, 0
	}

	var sep string
	if strings.Contains(string(match[0]), "-") {
		sep = "-"
	} else if strings.Contains(string(match[0]), ".") {
		sep = "."
	} else {
		sep = "/"
	}

	arr := strings.Split(string(match[0]), sep)
	if len(arr) != 1 && len(arr) != 3 {
		return 0, 0, 0
	}

	day, _ := strconv.ParseInt(arr[0], 10, 64)
	if len(arr) == 1 {
		return int(day), 0, 0
	}

	month, _ := strconv.ParseInt(arr[1], 10, 64)
	year, _ := strconv.ParseInt(arr[2], 10, 64)

	return int(day), int(month), int(year)
}

func calculateDate(exp [][]byte) (int, bool) {
	var result int
	var current int

	var add bool = true
	var formatType bool = true

	day, month, year := 0, 0, 0

	for i := 0; i < len(exp); i++ {
		val := string(exp[i])

		if val == "" || val == " + " || val == " - " {
			if val == " - " {
				add = false
			} else {
				add = true
			}
			continue
		}

		if val == "now" {
			date := time.Now()
			day, month, year = date.Day(), int(date.Month()), date.Year()
		} else {
			day, month, year = parseDate([]byte(val))
		}

		if day == 0 && month == 0 && year == 0 {
			fmt.Printf("Invalid time format at '%s'\n", val)
		}

		if month == 0 || year == 0 {
			current = day * 86_400
			formatType = false
		} else {
			current = int(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Unix())
			formatType = true
		}

		if !add {
			result -= current
		} else {
			result += current
		}
	}

	return result, formatType
}

func formatDate(result int, formatType bool) string {
	timeResult := time.Unix(int64(result), 0).UTC()
	days := result / 86_400

	if formatType {
		return fmt.Sprintf("%d days", days)
	} else {
		return fmt.Sprintf("%d %s %d", timeResult.Day(), timeResult.Month().String(), timeResult.Year())
	}

}
