package api

import "net/http"

// Регистрация всех API-обработчиков
func Init() {
    http.HandleFunc("/api/nextdate", nextDayHandler)
}
