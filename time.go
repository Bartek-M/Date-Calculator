package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseTime(val []byte) int {
	re := regexp.MustCompile(`^\d+(:\d+){0,2}$`)
	match := re.FindAll(val, -1)

	if len(match) != 1 {
		return 0
	}

	arr := strings.Split(string(match[0]), ":")
	var seconds int = 0

	for i := 0; i < len(arr); i++ {
		num, _ := strconv.ParseInt(arr[i], 10, 64)
		seconds += int(num) * int(math.Pow(60.0, float64(len(arr)-1-i)))
	}

	return seconds
}

func calculateTime(exp [][]byte) (int, bool) {
	var result int
	var num int
	var add bool = true

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
			now := time.Now()
			num = (now.Hour() * 3600) + (now.Minute() * 60) + now.Second()
		} else {
			num = parseTime([]byte(val))
		}

		if num == 0 && val != "0" {
			fmt.Printf("Invalid time format at '%s'\n", val)
		}

		if !add {
			result -= num
		} else {
			result += num
		}
	}

	return result, false
}

func formatTime(result int, _ bool) string {
	hours := result / 3600
	result -= hours * 3600

	minutes := result / 60
	result -= minutes * 60

	seconds := result

	return fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds)
}
