package controller

import (
	"net/http"
	"fmt"
	"../dao"
)

/**
测试链接
 */
func TestController(w http.ResponseWriter, r *http.Request) {

	fmt.Println("-------------这是Value获取的--------------")
	id := r.PostFormValue("id")
	fmt.Println("用户名: ", id)

	dao.Select()


	fmt.Fprintf(w, "欢迎使用Go语言开发服务器!") //这个写入到w的是输出到客户端的
}
