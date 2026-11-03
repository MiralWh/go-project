package works

import (
	"fmt"
	"go-project/app/helper"
	"math"
	"strings"
)

type PaymentProcessor interface {
	Process(amount float64) string
}

type CreditCard struct {
	Holder   string
	Number   string
	Bank     string
	ExpMonth int
	ExpYear  int
	Currency string
}

func (c CreditCard) Process(amount float64) string {
	if amount <= 0 {
		return fmt.Sprintf("CreditCard [%s]: отклонено — некорректная сумма %.2f", c.Number, amount)
	}
	// Маскируем номер ещё раз на всякий случай (вдруг в поле попал полный)
	masked := maskCard(c.Number)
	return fmt.Sprintf(
		"CreditCard [%s] (%s, %s): списано %.2f %s со счёта %s",
		masked, c.Bank, c.Holder, amount, c.Currency, masked,
	)
}

func maskCard(number string) string {
	digits := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, number)
	if len(digits) <= 4 {
		return "**** " + digits
	}
	return "**** **** **** " + digits[len(digits)-4:]
}

type CryptoWallet struct {
	Address string
	Network string
	Label   string
	Asset   string
}

func (w CryptoWallet) Process(amount float64) string {
	if amount <= 0 {
		return fmt.Sprintf("CryptoWallet [%s]: отклонено — некорректная сумма %.2f", w.Label, amount)
	}
	short := w.Address
	if len(short) > 12 {
		short = short[:6] + "..." + short[len(short)-4:]
	}
	return fmt.Sprintf(
		"CryptoWallet [%s] (%s, %s): отправлено %.2f %s на адрес %s",
		w.Label, w.Network, w.Asset, amount, w.Asset, short,
	)
}

var (
	_ PaymentProcessor = CreditCard{}
	_ PaymentProcessor = CryptoWallet{}
)

//------------------------------------------------------

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

var (
	_ Shape = Circle{}
	_ Shape = Rectangle{}
)

func printAreas(shapes []Shape) {
	for i, s := range shapes {
		fmt.Printf("%d) %-10T площадь = %.2f\n", i+1, s, s.Area())
	}
}

func Work6() {

	fmt.Println("Добро пожаловать на 6-e задание по языку \"Go\"")

	flag, task := helper.GetUserInputInt("Ввведите номер задачи 1 - 2")

	if flag {

		switch task {
		case 1:

			head := `Описание: Создайте интерфейс PaymentProcessor с методом Process(amount float64) string, который возвращает строку с результатом обработки платежа. Реализуйте два типа платежных систем: CreditCard и CryptoWallet, каждый из которых реализует интерфейс PaymentProcessor. В функции main создайте список платежных систем и вызовите метод Process для каждого, выводя результат на экран.
Требования:
    Интерфейс PaymentProcessor.
    Структуры CreditCard и CryptoWallet, реализующие интерфейс.
    В main создайте массив/слайс этих структур и вызовите их методы.
    P.S. Подумайте, какие поля могут быть у каждой структуры`
			fmt.Println(head)

			processors := []PaymentProcessor{
				CreditCard{
					Holder:   "Иван Петров",
					Number:   "1234 5678 9012 3456",
					Bank:     "Тинькофф",
					ExpMonth: 12,
					ExpYear:  2027,
					Currency: "RUB",
				},
				CryptoWallet{
					Address: "0xAbC1234567890dEf1234567890aBcDeF12345678",
					Network: "Ethereum",
					Label:   "Основной кошелёк",
					Asset:   "ETH",
				},
				CreditCard{
					Holder:   "Maria Sidorova",
					Number:   "9876 5432 1098 7654",
					Bank:     "Сбер",
					ExpMonth: 6,
					ExpYear:  2026,
					Currency: "RUB",
				},
				CryptoWallet{
					Address: "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
					Network: "Bitcoin",
					Label:   "Холодный кошелёк",
					Asset:   "BTC",
				},
			}

			amounts := []float64{1500.00, 0.025, 250, -10}

			for i, p := range processors {
				amount := amounts[i%len(amounts)]

				fmt.Println(p.Process(amount))
			}

			fmt.Println()

		case 2:

			head := `Описание: Создайте интерфейс Shape с методом Area() float64. Реализуйте структуры Circle и Rectangle, которые реализуют этот интерфейс. Напишите функцию, которая принимает слайс фигур ([]Shape) и выводит площадь каждой фигуры.
Требования:
    Интерфейс Shape.
    Структуры Circle (с радиусом) и Rectangle (с длиной и шириной).
    Функция для вывода площадей всех фигур в слайсе.`
			fmt.Println(head)

			shapes := []Shape{
				Circle{Radius: 5},
				Rectangle{Length: 4, Width: 6},
				Circle{Radius: 1.5},
				Rectangle{Length: 10, Width: 2.5},
			}

			printAreas(shapes)

			fmt.Println()

		default:
			fmt.Println("Неизвестная задача")
		}
	}
}
