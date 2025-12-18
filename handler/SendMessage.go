package handler

import (
	"encoding/json"
	"fakechat/db"
	"net/http"
)

type MessageRequest struct {
	Content string `json:"content"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从 cookie 获取用户名
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "not logged in", http.StatusUnauthorized)
		return
	}
	sender := cookie.Value

	var req MessageRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Content == "" {
		http.Error(w, "invalid content", http.StatusBadRequest)
		return
	}

	err = db.CreateMessage(sender, req.Content)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("message sent"))
}
