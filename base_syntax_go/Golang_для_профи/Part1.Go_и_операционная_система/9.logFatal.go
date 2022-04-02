// Пример использования log.Fatal().
// Функция log.Fatal() используется тогда, когда происходит что-то действительно очень плохое и надо просто поскорее выйти из программы, предварительно отправив сообщение о сложной ситуации
package main 

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}

// При использовании log.Fatal() программа завершается на том месте, где была вызвана функция log.Fatal()