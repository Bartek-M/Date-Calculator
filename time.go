package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func parseTime(val []byte) int {
	re := regexp.MustCompile(`^\d+(:\d+){0,2}$`)
	match := re.FindAll(val, -1)

	if len(match) != 1 {
		return 0
	}

	arr := strings.Split(string(match[0]), ":")
	var time int = 0

	for i := 0; i < len(arr); i++ {
		num, _ := strconv.ParseInt(arr[i], 10, 64)
		time += int(num) * int(math.Pow(60.0, float64(len(arr)-1-i)))
	}

	return time
}

func calculateTime(exp [][]byte) int {
	var result int = 0
	var add bool = true

	for i := 0; i < len(exp); i++ {
		val := string(exp[i])

		if val == "" {
			continue
		}

		if val == " + " {
			add = true
			continue
		} else if val == " - " {
			add = false
			continue
		}

		var num int = parseTime([]byte(val))

		if num == 0 && val != "0" {
			fmt.Printf("Invalid time format at '%s'\n", val)
		}

		if !add {
			result -= num
		} else {
			result += num
		}
	}

	return result
}

func formatTime(result int) string {
	hours := result / 3600
	result -= hours * 3600

	minutes := result / 60
	result -= minutes * 60

	seconds := result

	return fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds)
}