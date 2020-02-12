package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Robot is the model for our db robot
type Robot struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Type      string `json:"type" db:"type"`
	Dangerous bool   `json:"dangerous" db:"dangerous"`
}

func handleRobots(c *gin.Context) {
	db := c.Value("db").(*sqlx.DB)
	robots := []Robot{}
	err := db.Select(&robots, "SELECT * FROM robot")
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, robots)
}
