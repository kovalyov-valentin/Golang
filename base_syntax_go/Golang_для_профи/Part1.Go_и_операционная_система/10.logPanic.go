// Функция log.Panic()
/* Возможный ситуации, когда программа завершается сбоем и вы хотите получить как больше информации о нем.
В такие моменты одним из решений станет использование log.Panic() - это разновидность функций журналирования. */

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
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
// В итоге при исполнении программы мы видим, что log.Panic() выводит дополнительную низкоуровневую информацию.
// Как и log.Fatal(), функция log.Panic() добавит запись в соответствующий файд журнала и немедленно прекратит работу программы.