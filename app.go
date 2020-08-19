package main

import (
	"net/http"

	"github.com/bin16/Reversi/eventutil"
	"github.com/bin16/Reversi/reversic"
	"github.com/bin16/Reversi/userc"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/_debug", func(c *gin.Context) {
		c.String(http.StatusOK, "username="+userc.Username(c))
	})
	app.POST("/user.login", userc.UserLogin)
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
