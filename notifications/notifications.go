// Учебные уведомления: печать в консоль вместо почты и SMS.
package notifications

import "fmt"

// Любой способ отправить сообщение читателю.
type Notifier interface {
	Notify(readerName, message string) error
}

// «Почта».
type EmailNotifier struct {
	Email string
}

// Печатает сообщение как письмо.
func (n *EmailNotifier) Notify(readerName, message string) error {
	fmt.Printf("[EMAIL → %s] %s: %s\n", n.Email, readerName, message)
	return nil
}

// «СМС».
type SMSNotifier struct {
	Phone string
}

// Печатает сообщение как СМС.
func (n *SMSNotifier) Notify(readerName, message string) error {
	fmt.Printf("[SMS → %s] %s: %s\n", n.Phone, readerName, message)
	return nil
}
