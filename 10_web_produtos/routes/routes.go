package routes

import (
	"net/http"
	"web_produtos/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)

}
