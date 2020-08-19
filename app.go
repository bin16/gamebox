package main

import (
	"fmt"
	"net/http"

	"github.com/bin16/Reversi/eventutil"
	"github.com/bin16/Reversi/reversic"
	"github.com/bin16/Reversi/userc"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		name := userc.Username(c)
		nameOverride := c.Query("name")
		for _, n := range []string{nameOverride, name, "New Player"} {
			if n != "" {
				c.SetSameSite(http.SameSiteLaxMode)
				c.SetCookie("_u_", n, 99999999, "/", "", false, false)
				c.String(http.StatusOK, fmt.Sprintf("Ok, %s", n))
				return
			}
		}
	})
	app.GET("/_debug", func(c *gin.Context) {
		c.String(http.StatusOK, "username="+userc.Username(c))
	})
	app.POST("/reversi.create", reversic.CreateGame)

	gr := app.Group("/reversi")
	gr.GET("/", reversic.Index)
	gr.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		if sid, _ := c.Cookie("_sid"); sid != "" {
			eventutil.UnSubscribe(id, sid)
		}
		c.Next()
	}, reversic.Game)
	gr.POST("/:id/game.stat", reversic.StatGame)
	gr.POST("/:id/game.join", reversic.JoinGame)
	gr.GET("/:id/game.events", reversic.GameEvents)
	gr.POST("/:id/game.play/:name", reversic.PlayGame)

	app.Run(":2222")
}
