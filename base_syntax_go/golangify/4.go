// 1.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Вы находитесь в темной пещере.")

	var command = "выйти наружу"
	var exit = strings.Contains(command, "наружу")

	fmt.Println("Вы покидаете пещеру:", exit)
}

// 2. Представьте, что на выходе из темной пещеры яркие лучи полуденного солнца заставляют вас зажмуриться.
// Как бы вы объявили булеву переменную под названием wearShades?
// На английском слово «shades» может использоваться в значении «солнечные очки».

package main

import (
	"fmt"
	"strings"
)
func main() {
	fmt.Println("Вы находитесь в темной пещере.")

	var command = "выйти наружу"
	var exit = strings.Contains(command, "наружу")

	fmt.Println("Вы покидаете пещеру::", exit)

	fmt.Println("На выходе из темной пещеры яркие лучи полуденного солнца заставляют вас зажмуриться")

	var wearShades = "надеть солнечные очки"
	var wear = strings.Contains(wearShades, "надеть")

	fmt.Println("Вы надеваете солнечные очки:", wear)
}

// 3. Пример
package main

import "fmt"

func main() {
	fmt.Println("На знаке снаружи написано 'Несовершеннолетним вход запрещен'.")

	var age = 41
	var minor = age < 18

	fmt.Printf("В возрасте %v, я совершеннолетний? %v\n", age, minor)
}

// 4. Что больше — «яблоко» или «банан»?
package main

import "fmt"

func main() {
	fmt.Println("aplle" > "banana")
	fmt.Println("яблоко" > "банан")
}

// 5. Пример оператора if
package main

import "fmt"

func main() {
	var command = "идти на восток"
	if command == "идти на восток" {
		fmt.Println("Вы направляетесь к горе.")
	} else if command == "зайти внутрь" {
		fmt.Println("Вы заходите в пещеру, где будете жить до конца своей жизни.")
	} else {
		fmt.Println("Пока не совсем понятно.")
	}
}

// 6. В рассматриваемой приключенческой игре есть несколько локаций.
// Напишите программу, которая использует if и else if для отображения описания каждой из трех локаций: пещеры, входа и горы.
package main

import "fmt"

func main() {

	var location = ""
	if location == "caves" {
		fmt.Println("В пещере вниз головой висят летучие мыши.")
	} else if location == "entrance" {
		fmt.Println("У входа стоят канделябры со свечами.")
	} else if location == "mountains" {
		fmt.Println("На верхушке горы одиноко развивается флаг")
	} else {
		fmt.Println("Выберите локацию")
	}
}

// 7. Правила определения високосного года таковы:
// 	1. Любой год, что делится без остатка на четыре, но не делится без остатка на 100;
// 	2. Или любой год, что делится без остатка на 400.
	// Следует помнить, что для получения остатка деления двух чисел используется оператор %
	// Ноль в результате является показателем того, что первое число делится без остатка на второе.
package main

import "fmt"

func main() {
	fmt.Println("На дворе 2100 год. Он високосный?")

	var year = 2000
	var leap = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)

	if leap {
		fmt.Println("Этот год високосный!")
	} else {
		fmt.Println("К сожалению, нет. Этот год не високосный")
	}
}

// 8. Пример использования оператора switch
package main

import "fmt"

func main() {
	fmt.Println("Здесь вход в пещеру и путь на восток")
	var command = "прочитать знак"

	switch command { // Сравнивает case с command
	case "идти на восток":
		fmt.Println("Вы направляетесь в горе.")
	case "зайти в пещеру","зайти внутрь": // Запятая разделяет список возможных значений
		fmt.Println("Вы находитесь в тускло освещенной пещере.")
	case "прочитать знак":
		fmt.Println("На знаке написано 'Посторонним вход запрещен'")
	default:
		fmt.Println("Пока не совсем понятно")
	}
}

// 9. Пример
package main

import "fmt"

func main() {
	var room = "озеро"

	switch { // Выражения для каждого случая
	case room == "пещера":
		fmt.Println("Вы находитесь в тускло освещенной пещере.")
	case room == "озеро":
		fmt.Println("Лед кажется достаточно крепким.")
		fallthrough // Переходит на следующий случай
	case room == "глубина":
		fmt.Println("Вода такая холодная, что сводит кости.")
	}
}

// 10. Пример работы с for
package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10 // Объявление и инициализация

	for count > 0 { // Условие
		fmt.Println(count)
		time.Sleep(time.Second)
		count-- // Обратный отсчет; в противном случае цикл будет длиться вечно
	}
	fmt.Println("Запуск!")
}

// 11. В следующем коде какой-то объект движется по окружности 360 градусов и останавливается в случайной точке.
package main

import (
	"fmt"
	"math/rand"
)
func main() {
	var degrees = 0

	for {
		fmt.Println(degrees)

		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

// 12. Не каждый запуск проходит по плану.
// Реализуйте обратный отсчет, где на каждую секунду приходится шанс 1 к 100, что ввиду определенных обстоятельств запуск прервется, и счетчик остановится.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10 

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Запуск!")
	} else {
		fmt.Println("Запуск отменяется.")
	}
}

// 13. Напишите программу для угадывания числа. 
// Заставьте компьютер выбирать случайные числа из промежутка, пока он не угадает задуманное вами число. 
// Данное число нужно заранее объявить в верхней части программы. 
// Пускай на экране отображается каждая догадка с пояснением, больше данное число или меньше задуманного вами варианта.
package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 42 

	for {
		var n = rand.Intn(100) + 1
		if n < number {
			fmt.Printf("%v слишком маленькое число.\n", n)
		} else if n > number {
			fmt.Printf("%v слишком большое число.\n", n)
		} else {
			fmt.Printf("Угадал! %v\n", n)
			break
		}
	}
}