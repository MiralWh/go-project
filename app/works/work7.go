package works

import (
	"fmt"
	"go-project/app/helper"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Горутина %d начала работу\n", id)

	time.Sleep(time.Duration(id) * 300 * time.Millisecond)

	fmt.Printf("Горутина %d завершила работу\n", id)
}

//----------------------------------------------------------

func producer(ch chan<- int, n int) {
	defer close(ch)

	for i := 1; i <= n; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int, done chan<- struct{}) {
	defer close(done)

	for v := range ch {
		fmt.Println(v)
	}
}

func Work7() {

	fmt.Println("Добро пожаловать на 7-e задание по языку \"Go\"")

	flag, task := helper.GetUserInputInt("Ввведите номер задачи 1 - 2")

	if flag {

		switch task {
		case 1:

			head := `Описание: Напишите программу, которая запускает три горутины. Каждая горутина должна выводить сообщение о выполнении своей задачи и завершаться. Основная функция должна дождаться завершения всех горутин перед завершением программы.
Требования:
    Используйте sync.WaitGroup для синхронизации.
    Каждая горутина выводит сообщение вида: "Горутина X начала работу" и "Горутина X завершила работу".
    В основном потоке дождитесь завершения всех горутин.`
			fmt.Println(head)

			var wg sync.WaitGroup

			wg.Add(3)

			for i := 1; i <= 3; i++ {
				go worker(i, &wg)
			}

			fmt.Println("Основной поток: ожидаю завершения всех горутин...")
			wg.Wait()

			fmt.Println()

		case 2:

			head := `Описание: Создайте программу, в которой одна горутина генерирует числа от 1 до 10 и отправляет их через канал. Другая горутина читает числа из канала и выводит их на экран. После отправки всех чисел закройте канал и завершите программу.
Требования:
    Используйте каналы для передачи данных.
    Горутина-генератор должна отправлять числа и закрывать канал после этого.
    Горутина-потребитель должна читать из канала и выводить числа, пока канал не закрыт.`
			fmt.Println(head)

			ch := make(chan int)
			done := make(chan struct{})

			go producer(ch, 10)
			go consumer(ch, done)

			<-done
			fmt.Println("Готово: все числа переданы и выведены")

			fmt.Println()

		default:
			fmt.Println("Неизвестная задача")
		}
	}
}
