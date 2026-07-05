package websrvr

import (
	"fmt"
    "log"
    "net/http"
    "path/filepath"

    "github.com/StanislavDem/go-final-project/tests"
	"github.com/StanislavDem/go-final-project/pkg/api"
)

func StartServer() {
	// регистрируем API-обработчики
    api.Init()
	
    // Директория для фронтенда
    webDir := filepath.Join("web")

    // FileServer возвращает index.html и вложенные файлы
    fs := http.FileServer(http.Dir(webDir))

    // StripPrefix для того чтобы "/" вело прямо в webDir
    http.Handle("/", http.StripPrefix("/", fs))

    addr := ":" + fmt.Sprintf("%d", tests.Port) // порт берём из settings.go
    log.Printf("Server started on %s\n", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatal(err)
    }
}