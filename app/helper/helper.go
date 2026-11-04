package helper

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

const empty string = ""
const replacement string = "Input integer number"

func GetRootPathCombine(append string) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	str, err := filepath.Abs(file)
	if err != nil {
		return ""
	}

	index := strings.Index(str, "\\app")
	remove := str[index:]

	result := strings.ReplaceAll(str, remove, append)
	return result
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.ReplaceAll(fmt.Sprint(a), " ", delim), "[]")
}

func GetUserInputInt(prompt string) (bool, int64) {

	line := GetUserInput([]string{prompt})

	if len(line) > 0 {

		value, error := strconv.ParseInt(line[0], 10, 64)

		if error != nil {
			fmt.Println("Error occurred:", error)
			return false, math.MinInt
		}

		return true, value
	}

	return false, math.MinInt
}

func GetUserInputString(prompt string) (bool, string) {

	line := GetUserInput([]string{prompt})

	if len(line) > 0 {

		value := line[0]

		if utf8.RuneCountInString(value) == 0 {
			fmt.Println("Error occurred:", errors.New("invalid string"))
			return false, ""
		}

		return true, value
	}

	return false, ""
}

func GetUserInput(prompt []string) []string {

	size := len(prompt)
	reader := bufio.NewReader(os.Stdin)
	result := make([]string, size)

	for i := range size {

		line := prompt[i]

		fmt.Println(line)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		result[i] = input
	}

	return result
}

func GetIntValue(message string) int {
	var value int

	if message == empty {
		message = replacement
	} else {
		regexpr := regexp.MustCompile(`[\:\>]\s*$`)
		message = regexpr.ReplaceAllLiteralString(message, "")
	}

	message += ">"

	fmt.Println(message)
	fmt.Scan(&value)
	return value
}

func GetReverse(arr ...int) []int {

	len := len(arr)

	if len == 0 {
		return arr
	}

	j := len - 1

	for i := range len {

		k := arr[i]
		arr[i] = arr[j]
		arr[j] = k
		j--
	}

	return arr
}

func GetSorted(arr ...int) (lst []int, min int, max int, avr int) {

	len := len(arr)

	if len == 0 {
		return arr, 0, 0, 0
	}

	sum := 0

	for i := range len {

		sum += arr[i]

		for j := 0; j < (len - i - 1); j++ {

			if arr[j] > arr[j+1] {
				cur := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = cur
			}
		}
	}

	return arr, arr[0], arr[len-1], sum / len
}
