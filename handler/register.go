package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"fakechat/db"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	// 1. 确保是 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. 解析 JSON body
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("decode json error:", err)
		http.Error(w, "invalid request format", http.StatusBadRequest)
		return
	}

	// 3. 简单参数校验
	if req.Username == "" || req.Password == "" {
		http.Error(w, "username or password empty", http.StatusBadRequest)
		return
	}

	// 4. 使用 db 层检查用户是否存在
	exists, err := db.UserExists(req.Username)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "username already exists", http.StatusBadRequest)
		return
	}

	// 5. 使用 db 层创建用户
	err = db.CreateUser(req.Username, req.Password)
	if err != nil {
		log.Println("CreateUser error:", err)
		http.Error(w, "register failed", http.StatusInternalServerError)
		return
	}

	// 6. 成功返回
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("register success"))
}
