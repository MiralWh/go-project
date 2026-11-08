package works

import (
	"fmt"
	"go-project/app/helper"
	"strconv"
)

func Work1() {

	fmt.Println("Добро пожаловать на 1-e задание по языку \"Go\"")

	v1 := helper.GetIntValue("Введите 1-e целочисленное значение:")
	v2 := helper.GetIntValue("Введите 2-e целочисленное значение:")
	v3 := helper.GetIntValue("Введите 3-e целочисленное значение:")
	v4 := helper.GetIntValue("Введите 4-e целочисленное значение:")
	v5 := helper.GetIntValue("Введите 5-e целочисленное значение:")

	fmt.Println("Введены значения:")
	fmt.Printf("%d %d %d %d %d", v1, v2, v3, v4, v5)

	arr, min, max, avr := helper.GetSorted(v1, v2, v3, v4, v5)

	fmt.Println()
	fmt.Println("\"")
	fmt.Printf("Отсортированные элементы: %s", helper.ArrayToString(arr, ","))
	fmt.Println()
	fmt.Printf("Самое большое число: %s", strconv.Itoa(max))
	fmt.Println()
	fmt.Printf("Самое маленькое число %s", strconv.Itoa(min))
	fmt.Println()
	fmt.Printf("Среднее-арифметическое: %s", strconv.Itoa(avr))
	fmt.Println()
	fmt.Println("\"")
}
