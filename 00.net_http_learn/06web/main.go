package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	var company Company
	http.HandleFunc("/company", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			dec := json.NewDecoder(r.Body)
			err := dec.Decode(&company)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			enc := json.NewEncoder(w)
			err = enc.Encode(company)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":9091", nil)
}
