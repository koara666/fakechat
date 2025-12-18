package db

import (
	"database/sql"

	"log"

	"golang.org/x/crypto/bcrypt"
)

// 检查用户名是否已存在
func UserExists(username string) (bool, error) {
	row := DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// 创建用户（带 bcrypt hashing）
func CreateUser(username, password string) error {
	// bcrypt hash
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error:", err)
		return err
	}

	_, err = DB.Exec(
		"INSERT INTO users (username, password_hash) VALUES (?, ?)",
		username,
		string(hashed),
	)

	if err != nil {
		log.Println("CreateUser error:", err)
		return err
	}

	return nil
}

// 校验密码
func CheckPassword(username, password string) (bool, error) {
	var storedHash string
	err := DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&storedHash)

	if err == sql.ErrNoRows {
		return false, nil // 用户不存在
	}
	if err != nil {
		log.Println("CheckPassword error:", err)
		return false, err
	}

	// bcrypt 比较
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return false, nil // 密码不匹配
	}

	return true, nil
}
