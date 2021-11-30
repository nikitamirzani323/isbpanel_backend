package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nikitamirzani323/isbpanel_backend/db"
	"github.com/nikitamirzani323/isbpanel_backend/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":7072"))
}
