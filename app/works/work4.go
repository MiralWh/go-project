package works

import (
	"errors"
	"fmt"
	"go-project/app/helper"
)

func Work4() {

	fmt.Println("Добро пожаловать на 4-e задание по языку \"Go\"")

	flag, task := helper.GetUserInputInt("Ввведите номер задачи 1 - 4")

	if flag {

		switch task {
		case 1:

			head := `Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b. Если b равно нулю, возвращайте ошибку. В основной программе вызовите эту функцию и обработайте возможную ошибку.`
			fmt.Println(head)

			cases := []struct{ a, b float64 }{
				{10, 2},
				{-7, 4},
				{5, 0},
				{0, 0},
			}

			for _, c := range cases {
				result, err := divide(c.a, c.b)
				if err != nil {
					fmt.Printf("divide(%.2f, %.2f) -> ошибка: %v\n", c.a, c.b, err)
					continue
				}
				fmt.Printf("divide(%.2f, %.2f) = %.4f\n", c.a, c.b, result)
			}

			fmt.Println()

		case 2:

			head := `Функция высшего порядка: передача функции как аргумента
Создайте функцию applyOperation(a, b int, op func(int, int) int) int, которая применяет переданную функцию op к числам a и b.
Создайте несколько функций-операций: сложение, вычитание, умножение.
В основной программе вызовите applyOperation с разными операциями и выведите результаты.`
			fmt.Println(head)

			a, b := 12, 4

			fmt.Printf("applyOperation(%d, %d, add) = %d\n", a, b, applyOperation(a, b, add))
			fmt.Printf("applyOperation(%d, %d, sub) = %d\n", a, b, applyOperation(a, b, sub))
			fmt.Printf("applyOperation(%d, %d, mul) = %d\n", a, b, applyOperation(a, b, mul))

			mod := func(x, y int) int { return x % y }
			fmt.Printf("applyOperation(%d, %d, mod) = %d\n", a, b, applyOperation(a, b, mod))

			fmt.Printf("applyOperation(%d, %d, func(x,y) {x*y+x}) = %d\n",
				a, b, applyOperation(a, b, func(x, y int) int { return x*y + x }))

			fmt.Println()

		case 3:

			head := `Напишите программу, которая:
1) Создаёт одну анонимную функцию, которая при каждом вызове увеличивает внутренний счётчик на 1 и возвращает его значение.
2) Вызывает эту функцию несколько раз и выводит результат.
Пример:
// Вызовы функции
fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
fmt.Println(counter()) // 3`

			fmt.Println(head)

			count := 0
			counter := func() int {
				count++
				return count
			}

			fmt.Println(counter()) // 1
			fmt.Println(counter()) // 2
			fmt.Println(counter()) // 3

			fmt.Println()

		case 4:

			head := `Создайте функцию sumAll, которая принимает произвольное количество целых чисел и возвращает их сумму.
Пример использования:
fmt.Println(sumAll(1, 2, 3)) // 6
fmt.Println(sumAll(10, -2, 4, 7)) // 19`
			fmt.Println(head)

			fmt.Println(sumAll(1, 2, 3))      // 6
			fmt.Println(sumAll(10, -2, 4, 7)) // 19
			fmt.Println(sumAll())             // 0
			fmt.Println(sumAll(42))           // 42

			values := []int{5, 10, 15, 20}
			fmt.Println(sumAll(values...)) // 50

			fmt.Println()

		default:
			fmt.Println("Неизвестная задача")
		}
	}
}

var ErrorDividedByZero = errors.New("ошибка деления на ноль")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrorDividedByZero
	}
	return a / b, nil
}

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func sumAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
