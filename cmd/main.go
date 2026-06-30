package main

import (
	"log"

	"github.com/StanislavDem/go-final-project/websrvr"
	"github.com/StanislavDem/go-final-project/pkg/db"
)

func main() {
	    if err := db.Init("dataBase/scheduler.db"); err != nil { // проверка на наличие scheduler.db
        log.Fatalf("failed to init db: %v", err)
    }
	
    websrvr.StartServer()
}