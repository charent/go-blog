package main

import (
	"go-blog/router"
	"go-blog/utils/log"
)

func main() {
	log.Info.Println("server start")
	router.InitRouter()
}
