// 5. Пример использования оператора switch
package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Сегодня понедельник.")

	case time.Tuesday:
        fmt.Println("Сегодня вторник.")

    case time.Wednesday:
        fmt.Println("Сегодня среда.")

    case time.Thursday:
        fmt.Println("Сегодня четверг.")

    case time.Friday:
        fmt.Println("Сегодня пятница.")

    case time.Saturday:
        fmt.Println("Сегодня суббота.")

    case time.Sunday:
        fmt.Println("Сегодня воскресенье.")
	}
}

// 2. Пример того как поместить несколько выражений в одном кейсе
package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
		case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
			fmt.Println("будний день")
		case time.Saturday, time.Sunday:
			fmt.Println("выходные дни")
	}
}

// 3. Пример использования default
package main 

import "fmt" 

func main() {
	var size = "XXL"

	switch size {
	
	case "XXS":
        fmt.Println("очень очень маленький")
 
    case "XS":
        fmt.Println("очень маленький")
 
    case "S":
        fmt.Println("маленький")
 
    case "M":
        fmt.Println("средний")
 
    case "L":
        fmt.Println("большой")
 
    case "XL":
        fmt.Println("очень большой")
 
    case "XXL":
        fmt.Println("очень очень большой")
 
    default:
        fmt.Println("неизвестно")
	}
}

// 4. Пример необязательного оператора
package main 

import "fmt"

func main() {
	switch num := 6; num % 2 == 0 {
	case true:
		fmt.Println("even value")
	case false: 
		fmt.Println("odd value")
	}
}

// 5. Пример использования оператора break
package main 

import "fmt" 

func main() {
	w := "a b c \td\nefg hi"

	for _, e := range w {
		
		switch e {
		case ' ', '\t', '\n':
			break
		default:
			fmt.Printf("%c\n", e)
		}
	}
}

// 6. Пример использования switch без выражения
package main 

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	switch {
	case now.Hour() < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}
}

// 7. Пример с использованием ключевого слова fallthrough
package main 

import "fmt"

// A -> B -> C -> D -> E

func main() {

	nextspot := "B"

	fmt.Println("Stops ahead of us:")

	switch nextspot {
	
	case "A":
        fmt.Println("A")
        fallthrough
 
    case "B":
        fmt.Println("B")
        fallthrough
 
    case "C":
        fmt.Println("C")
        fallthrough
 
    case "D":
        fmt.Println("D")
        fallthrough
 
    case "E":
        fmt.Println("E")
	}
}

// 8. Использование type со switch
package main 

import "fmt" 

func main() {
	var data interface{}

	data = 112523652346.23463246345

	switch mytype := data.(type) {
		
	case string:
		fmt.Println("string")

	case bool:
		fmt.Println("boolean")

	case float64:
		fmt.Println("float64 type")

	case float32:
		fmt.Println("float32 type")

	case int:
		fmt.Println("int")
	default:
		fmt.Printf("%T", mytype)
	}
}