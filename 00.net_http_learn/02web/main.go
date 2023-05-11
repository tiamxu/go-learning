package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 创建web server的2种方法
// 1、http.ListenAndServe() 两个参数
// 2、http.Server 结构体，server.ListenAndServer()
type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}
func welcome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome"))
}
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web</title></head>
<body><h1>Hello World</h1></body>
</html>
`
	w.WriteHeader(301)

	w.Write([]byte(str))
}
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service,try netx door")
}
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	post := &Post{
		User:    "xxx",
		Threads: []string{"fist", "secend", "thread"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
func main() {
	my := helloHandler{}
	a := aboutHandler{}
	//1、
	// http.ListenAndServe(":9090", nil)
	//2、
	http.Handle("/hello", &my)
	http.Handle("/about", &a)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Home"))
		fmt.Fprintln(w, r.URL.RawQuery)

		url := r.URL
		// fmt.Fprintln(w, url)
		query := url.Query()
		fmt.Fprintln(w, query)
		id := query["id"]
		log.Println(id)
		name := query.Get("name")
		log.Println(name)

	})
	// http.HandleFunc("/welcome", welcome)
	http.Handle("/weclome", http.HandlerFunc(welcome))
	// http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, r.Header)
	// 	fmt.Fprintln(w, r.Header["User-Agent"])
	// 	fmt.Fprintln(w, r.Header.Get("User-Agent"))

	// })
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024)
		fmt.Fprintln(w, r.FormValue("first_name"))
		fmt.Fprintln(w, r.PostFormValue("first_name"))

	})

	http.HandleFunc("/fileupload", func(w http.ResponseWriter, r *http.Request) {
		// r.ParseMultipartForm(1024)
		// fileHeader := r.MultipartForm.File["upload"][0]
		// file, err := fileHeader.Open()
		file, _, err := r.FormFile("upload")
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}

	})
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/header", headerExample)
	http.HandleFunc("/json", jsonExample)

	server := http.Server{Addr: ":9091", Handler: nil}
	server.ListenAndServe()
}
