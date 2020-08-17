package userc

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Username(c *gin.Context) string {
	name, err := c.Cookie("_u_")
	if err != nil {
		log.Println(err)
	}

	return name
}
