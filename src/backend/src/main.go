package main

import "github.com/gin-gonic/gin"


// Form Data
type GreetForm struct {
	Name string `form:"name"`
}

type VoteForm struct {
  Age int `form:"age"`
}

func main() {
	server := gin.Default()
	server.GET("/api/hello", hello)
  server.GET("/api/vote", canVote)
	server.Run(":8081")
}

func hello(c *gin.Context) {
	var form GreetForm

	if c.ShouldBind(&form) == nil {
		c.JSON(200, gin.H{
			"message": "Hello " + form.Name,
		})
	}
}

func canVote(c *gin.Context) {
  var form VoteForm

  if c.ShouldBind(&form) == nil {
    if form.Age <= 0 {
      c.JSON(200, gin.H{
        "canVote": "Error: Invalid Age",
      })
    } else if form.Age > 18 {
      c.JSON(200, gin.H{
        "canVote": true,
      })
    } else {
      c.JSON(200, gin.H{
        "canVote": false,
      })
    }
  }
}

