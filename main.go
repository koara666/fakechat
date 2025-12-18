package main

import (
	"log"
	"net/http"

	"fakechat/db"
	"fakechat/handler"
)

func main() {
	// 1. 初始化数据库
	db.InitDB()

	// 2. 注册路由
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login) // 新增登录路由
	http.HandleFunc("/send", handler.SendMessage)
	http.HandleFunc("/messages", handler.GetMessages)
	// 3. 启动 HTTP 服务器
	log.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
