// 1. Пример области видимости
package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0

	for count < 10 { // Начало области видимости
		var num = rand.Intn(10) + 1
		fmt.Println(num)

		count++
	} // Конец области видимости  
}

// 2. В чем преимущество области видимости переменных?

Преимущество области видимости переменных заключается в том, что можно использовать одни и те же названия переменных, не переживая из-за возможных ошибок. Можно заниматься только теми переменными, которые находятся в данный момент в области видимости.

// 3. Что случится с переменной, когда она покинет область видимости? 
// Измените код предыдущего примера кода для получения доступа к num после цикла и посмотрите, что произойдет.

package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	
	for count < 10 {
		num = rand.Intn(10) + 1 // Переменная num больше не видна и не доступна
		fmt.Println(num)

		count++
	}
	var num = 0
}

// 4. Пример цикла for, что совмещает инициализацию, условие и последующий оператор
package main 

import "fmt" 

func main() {
	var count = 0
	for count = 10; count > 0; count -- {
		fmt.Println(count)
	}
	fmt.Println(count) // count остается в области видимости
}

// 5. Пример того, как переменная count объявляется и инициализируется как часть цикла for.
package main 

import "fmt" 

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	} // count больше не в области видимости
	// fmt.Println(count)
}

// 6. Пример того, как переменная num может использоваться в любом ответвлении оператора if.
package main 

import (
	"math/rand"
	"fmt" 
)
func main() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	} // num больше не в области видимости
}

// 7. Пример использования с оператором switch 
package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(11); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default: 
		fmt.Println("Random spaceline #", num)
	}
}

8. Как бы изменилась область видимости переменной num, если бы в предыдущих примерах не использовалось краткое объявление?
С помощью var объявить переменную сразу после ключевых слов if, switch и for невозможно.
В таком случает num требуется объявить перед операторами if/switch, чтобы переменная num осталась в области видимости за пределами оканчания if/switch.

// 9. Пример кода, который генерирует и отображает случайную дату — к примеру, дату вылета на Марс.
package main 

import (
	"fmt"
	"math/rand"
)
var era = "AD" // переменная era доступна через пакет

func main() {
	year := 2018 // переменные era и year находятся в области видимости 

	switch month := rand.Intn(12) + 1; month { // переменные era, year и month в области видимости
	case 2:
		day := rand.Intn(28) + 1 // новый день 
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)    
	} // month и day за пределами области видимости
} // year за пределами области видимости

// 10. Прошлый код после рефакторинга
package main 

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2: 
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}

// 11. Изменить вышепредставленную программу для обработки високосных годово.
// Код должен:
// 	Генерировать случайный год вместо постоянного использования 2018
// 	Для февраля присвойте daysInMonth на 29 для високосных годов, и 28 для всех остальных. Можете использовать оператор if вместо блока case
// 	Используйте цикл for для генерации и отображения 10 случайных дат.
package main 

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for count := 0; count < 10; count++{
		year := 2018 + rand.Intn(10)
		leap := year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
		month := rand.Intn(12) + 1

		daysInMonth := 31
		switch month {
		case 2:
			daysInMonth = 28
			if leap {
				daysInMonth = 29
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
	
}