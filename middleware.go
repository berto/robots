package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connectDB() gin.HandlerFunc {
	db, err := sqlx.Connect("postgres", "user=berto dbname=robots sslmode=disable")
	return func(c *gin.Context) {
		if err != nil {
			c.JSON(500, err.Error())
		}
		c.Set("db", db)
		c.Next()
	}
}
