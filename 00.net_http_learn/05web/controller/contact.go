package controller

import (
	"html/template"
	"net/http"
)

func contactHander(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "contact.html")
	t.ExecuteTemplate(w, "layout", "hello contact")
}

func registerContactRoutes() {
	http.HandleFunc("/contact", contactHander)

}
