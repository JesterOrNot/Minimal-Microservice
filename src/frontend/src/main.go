package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.StaticFile("/", "./static/index.html")
	server.Run()
}
