package main

import (
	"net/http"

	"github.com/clavinorach/crud-go/database"
	"github.com/clavinorach/crud-go/routes"
)

func main() {
	db := database.InitDatabase()
	
	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}