package controller

import (
	"net/http"
	"html/template"
	"fmt"
)

func HtmlController(w http.ResponseWriter, r *http.Request){
	fmt.Println("测试html跳转！！！")
	t, _ := template.ParseFiles("static/html/index.html")
	t.Execute(w, nil)
}
