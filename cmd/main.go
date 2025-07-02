package main

import (
	"fmt"

	"github.com/Jay-Chou118/mall/conf"
)

func main() {
	conf.Init()
	fmt.Println("Hello, 这是一个商城项目!")
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
