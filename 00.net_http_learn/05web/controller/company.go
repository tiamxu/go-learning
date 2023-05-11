package controller

import (
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func companiesHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "companies.html")
	t.ExecuteTemplate(w, "layout", nil)
}
func companyHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "company.html")
	patters, _ := regexp.Compile(`/companies/(\d+)`)
	matches := patters.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		companyID, _ := strconv.Atoi(matches[1])
		t.ExecuteTemplate(w, "layout", companyID)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
func registerCompanyRouters() {
	http.HandleFunc("/companies", companiesHandler)
	http.HandleFunc("/companies/", companyHandler)
}
