package works

import (
	"fmt"
	"go-project/app/helper"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("%q by %s, %d", b.Title, b.Author, b.Year)
}

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0.0
	for _, g := range s.Grades {
		sum += g
	}
	return sum / float64(len(s.Grades))
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Printf("Пополнение отклонено: сумма должна быть положительной (получено %.2f)\n", amount)
		return
	}
	b.Balance += amount
	fmt.Printf("Пополнено на %.2f. Баланс: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Printf("Снятие отклонено: сумма должна быть положительной (получено %.2f)\n", amount)
		return
	}
	if amount > b.Balance {
		fmt.Printf("Недостаточно средств: запрошено %.2f, доступно %.2f\n", amount, b.Balance)
		return
	}
	b.Balance -= amount
	fmt.Printf("Снято %.2f. Баланс: %.2f\n", amount, b.Balance)
}

func Work5() {

	fmt.Println("Добро пожаловать на 5-e задание по языку \"Go\"")

	flag, task := helper.GetUserInputInt("Ввведите номер задачи 1 - 3")

	if flag {

		switch task {
		case 1:

			head := `Создайте структуру Book, которая содержит поля: Title (строка), Author (строка), Year (целое число).
Добавьте метод GetInfo(), который возвращает строку с информацией о книге в формате: "Title" by Author, Year.`
			fmt.Println(head)

			book := Book{
				Title:  "Чужак в стране чужой",
				Author: "Роберт Хайнлайн",
				Year:   1962,
			}

			fmt.Println(book.GetInfo())

			fmt.Println()

		case 2:

			head := `Создайте структуру Student, которая содержит поля: имя (Name) и список оценок (Grades []float64).
Реализуйте метод AverageGrade() float64, который возвращает средний балл студента.
Что нужно сделать:
Объявить структуру и метод.
Создать студента с несколькими оценками и вывести его средний балл.`
			fmt.Println(head)

			student := Student{
				Name:   "Иван Иванов",
				Grades: []float64{4.5, 5.0, 3.8, 4.9, 5.0},
			}

			fmt.Printf("Студент: %s\n", student.Name)
			fmt.Printf("Оценки:  %v\n", student.Grades)
			fmt.Printf("Средний балл: %.2f\n", student.AverageGrade())

			empty := Student{Name: "Пётр Петров"}
			fmt.Printf("\nСтудент: %s\n", empty.Name)
			fmt.Printf("Средний балл: %.2f\n", empty.AverageGrade())

			fmt.Println()

		case 3:

			head := `Создайте структуру BankAccount с полями: Owner (владелец счета) и Balance (текущий баланс).
Реализуйте методы:
Deposit(amount float64), увеличивающий баланс.
Withdraw(amount float64), уменьшающий баланс, если хватает средств, иначе выводит сообщение о недостатке средств.
Что нужно сделать:
Объявить структуру и методы.
Создать счет, пополнить его, попытаться снять деньги и вывести итоговый баланс.`
			fmt.Println(head)

			acc := BankAccount{
				Owner:   "Иван Петров",
				Balance: 0,
			}

			fmt.Printf("Счёт открыт. Владелец: %s, баланс: %.2f\n\n", acc.Owner, acc.Balance)

			acc.Deposit(1000)
			acc.Deposit(500.50)
			fmt.Println()

			acc.Withdraw(200)
			acc.Withdraw(2000)
			acc.Withdraw(1300.50)
			fmt.Println()

			fmt.Printf("Итоговый баланс: %.2f\n", acc.Balance)

			fmt.Println()

		default:
			fmt.Println("Неизвестная задача")
		}
	}
}
