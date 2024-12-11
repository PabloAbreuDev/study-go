package main

import (
	"example/backend/db"
	"example/backend/internal/routes"
	"log"
)

func main() {

	db.ConnectMongoDB("mongodb://localhost:27017/finances-app")

	r := routes.SetupRoutes()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(r.Run(":8080"))

}
