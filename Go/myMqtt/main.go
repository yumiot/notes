package main

import (
	"fmt"
	"../myMqtt/config"
	"../myMqtt/dao"
	"time"
)

/**
启动函数
 */
func main() {
	var port = "8080"
	fmt.Println("服务器启动成功！","端口号：",port)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	config.WebConfig(port)
	dao.Init("root","123456","db_test")
}

