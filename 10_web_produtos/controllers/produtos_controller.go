package controllers

import (
	"html/template"
	"net/http"
	"web_produtos/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	p := models.GetProdutos()
	temp.ExecuteTemplate(w, "Index", p)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
