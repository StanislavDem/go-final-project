package main

import (
	"fmt"
    "log"
    "net/http"
    "path/filepath"

    "github.com/StanislavDem/go-final-project/tests"
)

func main() {
    // Директория для фронтенда
    webDir := filepath.Join("..", "web")

    // Возвращаем index.html при запросе "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.ServeFile(w, r, filepath.Join(webDir, "index.html"))
            return
        }
        // Для остальных файлов (css, js, изображения)
        http.ServeFile(w, r, filepath.Join(webDir, r.URL.Path))
    })

    addr := ":" + fmt.Sprintf("%d", tests.Port) // порт берём из settings.go
    log.Printf("Server started on %s\n", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatal(err)
    }
}