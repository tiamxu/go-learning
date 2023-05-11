package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.etcd.io/etcd/store"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}
type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	AuthorID  uint   `gorm:"not null"`
	Author    User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 显示博客的主页
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	// 显示文章列表
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// 显示单篇文章
}

func adminIndexHandler(w http.ResponseWriter, r *http.Request) {
	// 显示后台管理页面
}

func adminPostsHandler(w http.ResponseWriter, r *http.Request) {
	// 显示文章列表，并提供编辑和删除功能
}

func adminPostNewHandler(w http.ResponseWriter, r *http.Request) {
	// 显示新建文章的页面，并处理表单提交
}

func adminPostEditHandler(w http.ResponseWriter, r *http.Request) {
	// 显示编辑文章的页面，并处理表单提交
}

func adminPostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// 处理删除文章的请求
}
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session")
		if err != nil || session.Values["user_id"] == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		userID := session.Values["user_id"].(uint)
		var user User
		if err := db.First(&user, userID).Error; err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	}
}

func adminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(User)
		if user.Username != "admin" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/posts", postsHandler)
	r.HandleFunc("/posts/{id:[0-9]+}", postHandler)

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(adminMiddleware)

	admin.HandleFunc("/", adminIndexHandler)
	admin.HandleFunc("/posts", adminPostsHandler)
	admin.HandleFunc("/posts/new", adminPostNewHandler)
	http.ListenAndServe(":9999", lm)
}
