// Сохранение библиотеки в JSON-файл и чтение обратно.
package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vovavdovin123-wq/simple-library/domain"
	"github.com/vovavdovin123-wq/simple-library/library"
)

// Что лежит в файле: книги и читатели.
type storageData struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`
}

// Записать библиотеку в файл.
func SaveLibraryToJSON(filePath string, lib *library.Library) error {
	data := storageData{
		Books:   lib.Books,
		Readers: lib.Readers,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ") // структура → JSON с отступами
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	return os.WriteFile(filePath, jsonData, 0644) // записать файл
}

// Прочитать библиотеку из файла.
func LoadLibraryFromJSON(filePath string) (*library.Library, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data storageData
	if err := json.Unmarshal(jsonData, &data); err != nil { // JSON → структура
		return nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	lib := library.New()
	lib.Books = data.Books
	lib.Readers = data.Readers
	lib.UpdateIDs() // чтобы новые ID не совпали со старыми

	return lib, nil
}
