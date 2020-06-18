package user

import (
	"encoding/json"
	"log"
	"net/http"
)

// UserHandler 所有user模块的handler
func UserHandler() {
	http.HandleFunc("/login", loginController)
}

//登录
func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := LoginService(username, password)
	//把结构体转换为json数据
	b, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("用户结构体序列化失败，err= %v", err)
	}
	//设置响应内容为json
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
