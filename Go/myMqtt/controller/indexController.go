package controller

import (
	"net/http"
	"html/template"

	)

/**
测试链接
 */
func IndexController(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "欢迎使用Go语言开发服务器!") //这个写入到w的是输出到客户端的
	t, _ := template.ParseFiles("resources/templates/index.html")
	t.Execute(w, nil)
}
