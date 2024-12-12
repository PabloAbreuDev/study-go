package main

import (
	"example/backend/config"
	"example/backend/db"
	"example/backend/internal/routes"
	"log"
)

func main() {

	cfg := config.GetConfig()

	port := cfg.Port

	if port == "" {
		port = "8080"
	}

	db.ConnectMongoDB("mongodb://localhost:27017/finances-app")

	r := routes.SetupRoutes()
	log.Println("Server running on http://localhost:%s", port)
	log.Fatal(r.Run(":" + port))

}
