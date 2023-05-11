package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]User)

// 注册
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, ok := users[user.Username]; ok {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = string(hashedPassword)
	users[user.Username] = user
	w.WriteHeader(http.StatusCreated)
}

// 登陆
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	storedUser, ok := users[user.Username]
	if !ok {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "Login successful")

}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome,hello world.")
}
func Headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%v:%v\n", name, header)
		}
	}
}

type LoggingMiddleware struct {
	handler http.Handler
}

func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
}
func main() {
	mux := http.NewServeMux()
	// r := mux.NewRouter()
	// r.HandleFunc("/signup", SignupHandler).Methods("POST")

	mux.HandleFunc("/register", SignupHandler)
	mux.HandleFunc("/login", LoginHandler)

	lm := &LoggingMiddleware{handler: mux}
	// err := http.ListenAndServe(":9999", mux)
	// if err != nil {
	// 	fmt.Println("error")
	// }
	log.Fatal(http.ListenAndServe(":9999", lm))
}
