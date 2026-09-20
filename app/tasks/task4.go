package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"go-project/app/helperstructs"
)

type Item string

type Order struct {
	Id    int32
	Code  int32
	Items []Item
}

func getUserInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func createOrder(id int32, code int32, input []string) *Order {
	items := make([]Item, 3, 10)
	for i := range input {
		items[i] = Item(input[i])
	}
	return &Order{
		Id:    id,
		Code:  code,
		Items: items,
	}
}

func moveOrderToDone(order *Order) bool {
	if order != nil && order.Code != 1 {
		order.Code = 1
		return true
	}
	return false
}

func Task4() {

	reader := bufio.NewReader(os.Stdin)

	ordersDict := make(map[int32]*Order)
	ordersWork := make(map[int32]struct{})
	ordersDone := make(map[int32]struct{})
	var ordersCurr int32 = 0

	fmt.Println("---Приветсвуем в приложении кассы---")

	options :=
		`
	  0) вывести помощь
	  1) вывести все неготовые заказы
	  2) создать новый заказ
	  3) перевести заказ в "готовые"
	  4) выход
	`

	option := 0
	number := 0
	common := ""

	for {

		fmt.Println("Доступные команды:")
		fmt.Println(options)

		input := getUserInput(reader)

		fmt.Println("Вы ввели:", input)

		switch {

		case strings.Contains(input, "помощь"):
			option = 0
		case strings.Contains(input, "неготовые"):
			option = 1
		case strings.Contains(input, "новый"):
			{
				fmt.Println("Введите состав заказа:")
				common = getUserInput(reader)
				option = 2
			}
		case strings.Contains(input, "перевести"):
			{
				fmt.Println("Введите номер заказа:")
				common = getUserInput(reader)
				locals, failer := strconv.Atoi(common)
				if failer != nil {
					option = 0
				} else {
					number = locals
				}
				option = 3
			}
		case strings.Contains(input, "выход"):
			option = 4
		default:
			{
				fmt.Println("неопознанная команда:")
				fmt.Printf("> %s", input)
				fmt.Println()
				option = 0
			}
		}

		if option == 4 {
			break
		}

		if option != 0 {

			switch option {

			case 1:
				{
					fmt.Println("Невыполненные заказы:")

					for i := range ordersWork {
						i32 := int32(i)
						fmt.Printf("%#v", ordersDict[i32])
						fmt.Println()
					}

					fmt.Println("==========================")
				}

			case 2:
				{
					words := strings.Split(common, ",")

					for i := range words {
						words[i] = strings.TrimSpace(words[i])
					}

					ordersCurr++
					new := createOrder(ordersCurr, 0, words)
					ordersDict[ordersCurr] = new
					ordersWork[ordersCurr] = struct{}{}
					fmt.Println("Новый заказ:")
					fmt.Printf("%#v", *new)
					fmt.Println()
					fmt.Println("==========================")
				}

			case 3:
				{
					id := int32(number)
					var item = ordersDict[id]
					if !moveOrderToDone(item) {
						fmt.Println("Что-то пошло не так...")
					} else {

						delete(ordersWork, id)
						ordersDone[ordersCurr] = struct{}{}
						fmt.Printf("Заказ %#v успешно выполнен!", item)
						fmt.Println()

					}
					option = 0
				}
			}
		}
	}

	fmt.Println("Выход из программы...")
	fmt.Println("Нажмите любую кнопку...")
	fmt.Scanln()
}
