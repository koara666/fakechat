package handler

import (
	"encoding/json"
	"net/http"

	"fakechat/db"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "username or password empty", http.StatusBadRequest)
		return
	}

	// 这里调用 db 层校验
	ok, err := db.CheckPassword(req.Username, req.Password)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: req.Username,
		Path:  "/",
	})

	w.Write([]byte("login success"))
}
