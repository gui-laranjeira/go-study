package main

import (
	"fmt"
	"net/http"
	"web_produtos/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()

	fmt.Println("Server Listening at :8080")
	http.ListenAndServe(":8080", nil)
}
