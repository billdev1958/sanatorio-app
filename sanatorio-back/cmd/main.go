package main

import (
	"log"
	"sanatorioApp/internal/app"
)

func main() {

	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer application.DB.Close()
	if err := application.Run(); err != nil {
		log.Fatalf("failed tols run app: %v", err)
	}

}
