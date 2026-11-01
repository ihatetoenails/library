// Книга и читатель — данные, без меню и файлов.
package domain

import "fmt"

// Книга. json-тег — имя поля в файле.
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	IsIssued bool   `json:"is_issued"` // выдана ли
	ReaderID int    `json:"reader_id"` // кому выдана
}

// Читатель.
type Reader struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active"` // можно ли выдавать книги
}

// Как печатать книгу в консоли.
func (b *Book) String() string {
	return fmt.Sprintf("'%s' (%s, %d)", b.Title, b.Author, b.Year)
}

// Отметить книгу выданной. Ошибка, если она уже у кого-то.
func (b *Book) IssueBook(readerID int) error {
	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана читателю с ID %d", b.Title, b.ReaderID)
	}
	b.IsIssued = true
	b.ReaderID = readerID
	return nil
}

// Вернуть книгу на полку. Ошибка, если она и так не выдана.
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = 0
	return nil
}

// Как печатать читателя: имя и фамилия.
func (r *Reader) String() string {
	return fmt.Sprintf("%s %s", r.FirstName, r.LastName)
}

// Разрешить выдачу книг этому читателю.
func (r *Reader) Activate() {
	r.IsActive = true
}

// Запретить выдачу книг.
func (r *Reader) Deactivate() {
	r.IsActive = false
}
