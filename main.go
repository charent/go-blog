package main

import (
	"go-blog/router"
	"go-blog/utils/mylog"
)

func main() {
	mylog.Info.Println("server start")
	router.InitRouter()
}
