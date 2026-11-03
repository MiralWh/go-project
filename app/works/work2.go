package works

import (
	"fmt"
	"go-project/app/helper"
	"go-project/app/helperstrings"
)

func Work2() {

	fmt.Println("Добро пожаловать на 2-e задание по языку \"Go\"")

	flag, task := helper.GetUserInputInt("Ввведите номер задачи 1 - 4")

	if flag {

		switch task {
		case 1:

			flag, v1 := helper.GetUserInputString("Введите строку для подсчёта количества символов:")

			if !flag {
				fmt.Println("Что-то пошло не так")
			}

			size := helperstrings.GetStringSize(v1)

			fmt.Printf("Количество символов: %v", size)
			fmt.Println()

		case 2:

			flag, v2 := helper.GetUserInputString("Введите строку для подсчёта количества гласных букв (а, е, ё, и, о, у, ы, э, ю, я):")

			if !flag {
				fmt.Println("Что-то пошло не так")
			}

			flag, data := helperstrings.GetRuneCountInString(v2, []rune("аеёиоуыэюя"))

			if !flag {
				fmt.Println("Что-то пошло не так")
			}

			for i := range len(data) {

				item := data[i]

				fmt.Printf("Количество символов: %c :  %v", item.Char, item.Size)
				fmt.Println()
			}

		case 3:

			flag, v3 := helper.GetUserInputString("Введите строку для преобразования в заглавную первых букв каждого слова:")

			if !flag {
				fmt.Println("Что-то пошло не так")
			}

			flag, data := helperstrings.CapitalizeWords(v3)

			if !flag {
				fmt.Println("Что-то пошло не так")
			}

			fmt.Printf("%v -> %v", v3, data)
			fmt.Println()

		case 4:

			flag, v4 := helper.GetUserInputString("Введите строку для проверки соответствия круглых скобок:")

			if !flag {
				fmt.Println("Что-то пошло не так")
			}

			flag, open, close := helperstrings.СheckFormulaBrackets(v4)

			fmt.Printf("строка на вход: %v:", v4)
			fmt.Println()

			if flag {

				fmt.Printf("вывод: Скобки расставлены верно, %v открывающиеся, %v закрывающиеся", open, close)
				fmt.Println()

			} else {

				fmt.Printf("вывод: Скобки расставлены неправильно, %v открывающиеся, %v закрывающиеся", open, close)
				fmt.Println()
			}

		default:
			fmt.Println("Неизвестная задача")
		}
	}
}
