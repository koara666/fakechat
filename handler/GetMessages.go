package handler

import (
	"encoding/json"
	"fakechat/db"
	"net/http"
	"strconv"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	// 默认获取最近 50 条
	limit := 50
	q := r.URL.Query().Get("limit")
	if q != "" {
		if l, err := strconv.Atoi(q); err == nil {
			limit = l
		}
	}

	msgs, err := db.GetMessages(limit)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(msgs)
}
