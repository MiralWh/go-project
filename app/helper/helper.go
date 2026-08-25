package helper

import (
	"fmt"
	"regexp"
	"strings"
)

const empty string = ""
const replacement string = "Input integer number"

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.ReplaceAll(fmt.Sprint(a), " ", delim), "[]")
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
