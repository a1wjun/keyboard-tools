package main

import (
	"kle-app/web"
	"log"
	"os"
)

func main() {
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(cors)
	log.Fatal("Error", app.Serve())
}

