// Настройки приложения.
package config

import "errors"

// Взять порт из карты настроек. Ошибка, если ключа PORT нет.
func GetPort(cfg map[string]string) (string, error) {
	port, ok := cfg["PORT"]
	if !ok {
		return "", errors.New("ключ PORT отсутствует в конфигурации")
	}
	return port, nil
}
