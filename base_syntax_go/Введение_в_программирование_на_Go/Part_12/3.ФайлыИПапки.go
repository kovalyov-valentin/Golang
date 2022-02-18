// // Так мы можем создать файл:

// package main 

// import (
// 	"os"
// )

// func main() {
// 	file, err := os.Create("test.txt")
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()

// 	file.WriteString("Hello World")
// }

// // Для открытия файла в Го используется функция Open из пакета os.
// // Вот пример того, как прочитать файл и выести его содержимое на консоль.

// package main

// import (
// 	"fmt"
// 	"os"
// )
// func main() {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		// здесь перехватывается ошибка
// 		return
// 	}
// 	defer file.Close()

// 	// получить размер файла
// 	stat, err := file.Stat()
// 	if err != nil {
// 		return
// 	}
// 	// чтение файла
// 	bs := make([]byte, stat.Size())
// 	_, err = file.Read(bs)
// 	if err != nil {
// 		return
// 	}
// 	str := string(bs)
// 	fmt.Println(str)
// }

// // Чтение файлов является частым действием, вот самый короткий способ сделать это:

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	bs, err := ioutil.ReadFile("test.txt")
// 	if err != nil {
// 		return
// 	}
// 	str := string(bs)
// 	fmt.Println(str)
// }

// // Для того, чтобы получить содержимое каталога, мы используем все тот же os.Open.
// // Но вместо имени файла мы передаем ему путь к каталогу.
// // И затем вызывается функция Readdir

// package main

// import (
// 	"fmt"
// 	"os"
// )
// func main() {
// 	dir, err := os.Open(".")
// 	if err != nil {
// 		return
// 	}
// 	defer dir.Close()

// 	fileInfos, err := dir.Readdir(-1)
// 	if err != nil {
// 		return
// 	}
// 	for _, fi := range fileInfos {
// 		fmt.Println(fi.Name())
// 	}
// }

// Для того, чтобы рекурсивно обойти каталоги(прочитать содержимое текущего и всех вложенных каталогов),
// слeдует использовать функцию Walk, которая предоставляется пакетом path/filepath:

package main

import (
	"fmt"
	"os"
	"path/filepath"
)
func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}