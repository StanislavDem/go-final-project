package api

import "net/http"

// Регистрация маршрута
func taskHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        addTaskHandler(w, r)
    }
}
// Регистрация API-обработчиков
func Init() {
    http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
}