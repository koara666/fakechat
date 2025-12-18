package db

import (
	"database/sql"
	"errors"
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

// 创建用户
func CreateUser(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username or password empty")
	}

	_, err := DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

// 检查用户名与密码是否匹配
func CheckPassword(username, password string) (bool, error) {
	row := DB.QueryRow("SELECT password FROM users WHERE username = ?", username)

	var storedPassword string
	err := row.Scan(&storedPassword)

	if err == sql.ErrNoRows {
		return false, nil // 用户不存在
	}
	if err != nil {
		return false, err // 数据库查询错误
	}

	// 简单比较（未加密）
	return storedPassword == password, nil
}
