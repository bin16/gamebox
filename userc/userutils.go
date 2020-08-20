package userc

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Username(c *gin.Context) string {
	name, err := c.Cookie("_u_")
	if err != nil {
		log.Println(err)
	}

	return name
}

func SetUsername(c *gin.Context, name string) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("_u_", name, 999999999, "/", "", false, false)
}
