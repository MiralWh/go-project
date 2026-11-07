package works

import (
	"fmt"
	"go-project/app/helper"
)

func Work9() {

	fmt.Println("Добро пожаловать на 9-e задание по языку \"Go\"")

	flag, task := helper.GetUserInputInt("Ввведите номер задачи 1 - 2")

	if flag {

		switch task {
		case 1:

			head := `Вам необходимо написать тесты для задач 1 и 4 к теме 8.
Задача 1
Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b. Если b равно нулю, возвращайте ошибку. В основной программе вызовите эту функцию и обработайте возможную ошибку`

			fmt.Println(head)

			fmt.Println()

			data := `
# Все тесты пакета
# Запускать из папки app
go test -v ./works/

# Покрытие
go test -cover ./works/

# Только основной
go test -v -run '^TestDivide$' ./works/`

			fmt.Println(data)

			fmt.Println()

		case 2:

			head := `Вам необходимо написать тесты для задач 1 и 4 к теме 8.
Задача 4
Создайте функцию sumAll, которая принимает произвольное количество целых чисел и возвращает их сумму.
Пример использования:
    fmt.Println(sumAll(1, 2, 3)) // 6
    fmt.Println(sumAll(10, -2, 4, 7)) // 19`

			fmt.Println(head)

			fmt.Println()

			data := `
# Все тесты пакета
# Запускать из папки app
go test -v ./works/

# Покрытие
go test -cover ./works/

# Только основной
go test -v -run '^TestSumAll$' ./works/`

			fmt.Println(data)

			fmt.Println()

		default:
			fmt.Println("Неизвестная задача")
		}
	}
}
