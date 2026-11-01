// Запуск консольной библиотеки.
package main

// Печать, лог, файлы, логика библиотеки, чтение и запись JSON.
import (
	"fmt"
	"log"
	"os"

	"github.com/vovavdovin123-wq/simple-library/library"
	"github.com/vovavdovin123-wq/simple-library/storage"
)

// Загрузить books.json или создать пустую библиотеку, затем открыть меню.
func main() {
	const dbFile = "books.json"

	myLibrary, err := storage.LoadLibraryFromJSON(dbFile)
	if err != nil {
		if os.IsNotExist(err) { // файла ещё нет
			fmt.Println("Файл данных не найден, создана новая библиотека.")
			myLibrary = library.New()
		} else {
			log.Fatalf("Ошибка при загрузке библиотеки: %v", err)
		}
	} else {
		fmt.Println("Библиотека успешно загружена из файла.")
	}

	Run(myLibrary, dbFile) // меню
}
