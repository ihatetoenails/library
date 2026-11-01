// Запуск HTTP-сервера.
package main

// JSON, печать, лог, HTTP, время.
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Ответ /healthcheck: жив ли сервис.
type healthCheckResponse struct {
	Status      string `json:"status"`
	ProjectName string `json:"project_name"`
}

// Ответ /api/v1/info: версия и время сервера.
type infoResponse struct {
	Version    string `json:"version"`
	ServerTime string `json:"server_time"`
}

// Отдаёт JSON «сервис доступен».
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := healthCheckResponse{
		Status:      "available",
		ProjectName: "Library API",
	}

	jsonData, err := json.Marshal(response) // структура → JSON
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Content-Type", "application/json") // тело — JSON
	w.WriteHeader(http.StatusOK)                       // 200
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

// Отдаёт версию API и текущее время.
func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := infoResponse{
		Version:    "1.0.0",
		ServerTime: time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

// Ответ 404: такой страницы нет.
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("This page does not exist")); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

// Маршруты и запуск сервера на порту 8080.
func main() {
	http.HandleFunc("/healthcheck", healthCheckHandler) // адрес → функция
	http.HandleFunc("/api/v1/info", infoHandler)
	http.HandleFunc("/notfound", notFoundHandler)

	addr := ":8080"
	fmt.Println("Сервер запущен на http://localhost" + addr)
	fmt.Println("Проверка:   http://localhost" + addr + "/healthcheck")
	fmt.Println("Информация: http://localhost" + addr + "/api/v1/info")
	fmt.Println("404:        http://localhost" + addr + "/notfound")

	// Слушать запросы. nil — маршруты, заданные выше.
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
