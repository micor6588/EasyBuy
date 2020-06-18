package main

import (
	"EasyBuy/user"
	"html/template"
	"net/http"
)

//显示登录页面
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)
}
func main() {
	s := http.Server{Addr: ":8080"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	//调用所有user模块的handler
	user.UserHandler()
	s.ListenAndServe()
}
