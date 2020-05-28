package main

import "github.com/gin-gonic/gin"


// Form Data
type Form struct {
	Name string `form:"name"`
}

func main() {
	server := gin.Default()
	server.GET("/api/hello", hello)
	server.Run(":8081")
}

func hello(c *gin.Context) {
	var form Form

	if c.ShouldBind(&form) == nil {
		c.JSON(200, gin.H{
			"message": "Hello " + form.Name,
		})
	}
}
