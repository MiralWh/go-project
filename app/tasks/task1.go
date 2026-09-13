package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Работа по печати чека
func Task1() {
	reader := bufio.NewReader(os.Stdin)

	data := [3][2]string{
		{"", ""},
		{"", ""},
		{"", ""},
	}

	count := 0.0

	for i := range 3 {

		fmt.Println("enter command:")

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		fmt.Println("Вы ввели:", input)

		words := strings.Split(input, " ")

		if len(words) == 2 {
			data[i][0] = words[0]
			data[i][1] = words[1]

			val, _ := strconv.ParseFloat(words[1], 64)
			count += val
		}
	}

	fmt.Println("Есть ли карта постоянного покупателя? <Y/N>")
	card, _ := reader.ReadString('\n')

	card = strings.TrimSpace(card)

	hasDiscount := card == "Y"

	discount := 0

	total := count

	cell := count

	if hasDiscount {
		discount = 5

		cell *= 0.05

		total = total - cell
	}

	const width = 40
	line := strings.Repeat("=", width)
	dash := strings.Repeat("-", width)

	var b strings.Builder

	b.WriteString(line)
	b.WriteString("\n")
	b.WriteString(center("ОАО Golang", width))
	b.WriteString("\n")
	b.WriteString(line)
	b.WriteString("\n")

	for i := range 3 {
		fmt.Fprintf(&b, " %s\n", data[i][0])
		fmt.Fprintf(&b, "  %d x %s\n", i+1, data[i][1])
		fmt.Fprintf(&b, "  Скидка: %v %s\n", discount, "%")
		fmt.Fprintf(&b, "  Сумма: %.4f р\n", total)
	}

	b.WriteString(dash)
	b.WriteString("\n")

	//TODO: доработать вывод итогов
	fmt.Fprintf(&b, "Итого без скидки: %v\n", count)
	fmt.Fprintf(&b, "Скидка: %v\n", discount)
	fmt.Fprintf(&b, "Итог:  %.4f\n", total)

	b.WriteString(dash)
	b.WriteString("\n")
	b.WriteString(center(" Спасибо за покупку", width))
	b.WriteString("\n")
	b.WriteString(dash)
	b.WriteString("\n")

	fmt.Println(b.String())

	fmt.Println("Press any key to exit")
	fmt.Scanln()
}

func center(s string, width int) string {
	n := utf8.RuneCountInString(s)
	if n >= width {
		return s
	}
	left := (width - n) / 2
	right := width - n - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
