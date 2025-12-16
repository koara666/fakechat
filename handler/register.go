package handler

import (
	"log"
	"net/http"

	"fakechat/db"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// 1. 只允许 POST 请求
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST allowed"))
		return
	}

	// 2. 从表单中取数据
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username or password empty"))
		return
	}

	// 3. 插入数据库
	_, err := db.DB.Exec(
		"INSERT INTO users (username, password_hash) VALUES (?, ?)",
		username,
		password,
	)
	if err != nil {
		log.Println("insert user error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("register failed"))
		return
	}

	// 4. 返回成功
	w.Write([]byte("register success"))
}
