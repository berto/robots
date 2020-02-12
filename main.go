package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(connectDB())
	r.GET("/robots", getRobotHandler)
	r.POST("/robots", addRobotHandler)
	r.DELETE("/robots/:id", deleteRobotHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
