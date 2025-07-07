package main

import (
	"fmt"
	"github.com/Jay-Chou118/mall/routes"
	"net/http"

	"github.com/Jay-Chou118/mall/conf"
)

func main() {
	conf.Init()
	fmt.Println("Hello, 这是一个商城项目!")
	r := routes.NewRouter()
	port := fmt.Sprintf(":%s", conf.HttpPort)
	if err := r.Run(port); err != nil && err != http.ErrServerClosed {
		fmt.Printf("启动服务失败: %v\n", err)
	}
}
