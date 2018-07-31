package config

import (
	"net/http"
	"log"

	"../../myMqtt/controller"
)

/**
web的配置目录
 */
func WebConfig(port string) {

	// 设置静态目录
	//fsh := http.FileServer(http.Dir("/home/make/go/src/awesomeProject/src/myMqtt/resources"))
	fsh := http.FileServer(http.Dir("./resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fsh))

	//配置访问的路由
	http.HandleFunc("/", controller.IndexController)
	http.HandleFunc("/json",controller.JsonController)
	http.HandleFunc("/html",controller.HtmlController)
	http.HandleFunc("/login",controller.LoginController)
	http.HandleFunc("/u/login",controller.LoginControllerU)
	http.HandleFunc("/registered",controller.RegisteredController)
	http.HandleFunc("/t",controller.TestController)

	//监听端口
	webPort(port)
}


//设置监听的端口
func webPort(port string) {
	port = ":" + port
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
