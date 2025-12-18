package db

import (
	"log"
)

type Message struct {
	ID        int    `json:"id"`
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// 存消息
func CreateMessage(sender, content string) error {
	_, err := DB.Exec(
		"INSERT INTO messages (sender, content) VALUES (?, ?)",
		sender, content,
	)
	if err != nil {
		log.Println("CreateMessage error:", err)
		return err
	}
	return nil
}

// 获取最近 n 条消息
func GetMessages(limit int) ([]Message, error) {
	rows, err := DB.Query(
		"SELECT id, sender, content, created_at FROM messages ORDER BY id DESC LIMIT ?",
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Sender, &m.Content, &m.CreatedAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, m)
	}

	// 返回按时间升序
	for i, j := 0, len(msgs)-1; i < j; i, j = i+1, j-1 {
		msgs[i], msgs[j] = msgs[j], msgs[i]
	}

	return msgs, nil
}
