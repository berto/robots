package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// Robot is the model for our db robot
type Robot struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Type      string `json:"type" db:"type"`
	Dangerous bool   `json:"dangerous" db:"dangerous"`
}

// Validate checks to make sure all fields are valid
func (r Robot) Validate() bool {
	if r.Name == "" && len(r.Name) < 1000 {
		return false
	}
	if r.Type == "" {
		return false
	}
	return true
}

func getRobotHandler(c *gin.Context) {
	db := c.Value("db").(*sqlx.DB)
	robots := []Robot{}
	err := db.Select(&robots, "SELECT * FROM robot")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, robots)
}

func addRobotHandler(c *gin.Context) {
	// grab the robot from the body of the request
	robot := Robot{}
	err := c.BindJSON(&robot)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	id := uuid.NewV4()
	robot.ID = id.String()
	// validate robot information
	if ok := robot.Validate(); !ok {
		c.JSON(500, "invalid robot information")
		return
	}
	// insert robot into db
	db := c.Value("db").(*sqlx.DB)
	_, err = db.NamedExec(`
		INSERT INTO
			robot (id, name, type, dangerous)
		VALUES (:id,:name,:type,:dangerous);
		`,
		&robot,
	)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	// respond with a 200 and new id
	c.JSON(200, robot)
}

func deleteRobotHandler(c *gin.Context) {
	id := c.Param("id")

	db := c.Value("db").(*sqlx.DB)
	result, err := db.Exec(`
		DELETE FROM robot
		WHERE id = $1;
	`,
		id,
	)
	if err != nil {
		fmt.Println("INSIDE ERR", result)
		c.JSON(500, err.Error())
		return
	}
	rows, err := result.RowsAffected()
	if rows < 1 || err != nil {
		c.JSON(500, "unable to delete")
		return
	}
	c.JSON(200, id)
}
