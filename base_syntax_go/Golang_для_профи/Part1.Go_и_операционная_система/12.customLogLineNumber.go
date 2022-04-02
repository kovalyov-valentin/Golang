// Вывод номеров строк в записях журнала
// Пример программы, которая укажет в журнальном файле номер строки исходного файла, который выполнил инструкцию, сделавшую запись в журнал.
package main 

import (
	"log"
	"fmt"
	"os"
)

var LOGFILE = "/tmp/mGO.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "customLogLineNumber", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile) // Всю магию выполняет данный оператор, который кроме log.LstdFlags также активизирует флаг log.Lshortfile, который добавляет в строку записи журнала полное имя файла, а также номер строки оператора, который создал эту запись.
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
// Выполнение данной программы не выводит данные на экран. Но после двух запусков файла customLogLineNumber.go изменится содержимое журнала /tmp/mGo.log
// Ключевое слово defer поможет генерировать более красивые сообщения журнала. Но об этом позже.