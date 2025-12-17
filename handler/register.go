package handler

import (
	"database/sql"
	"log"
	"net/http"

	"fakechat/db"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// 1. 只允许 POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST allowed"))
		return
	}

	// 2. 取参数
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username or password empty"))
		return
	}

	// 3. 检查用户名是否存在
	var id int
	err := db.DB.QueryRow(
		"SELECT id FROM users WHERE username = ?",
		username,
	).Scan(&id)

	if err == nil {
		// 能查到，说明用户名已存在
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username already exists"))
		return
	}

	if err != sql.ErrNoRows {
		// 发生了其他数据库错误
		log.Println("query user error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
		return
	}

	// 4. 插入新用户
	_, err = db.DB.Exec(
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

	// 5. 成功
	w.Write([]byte("register success"))
}
