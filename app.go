package main

import (
	"net/http"

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
	gr.GET("/:id", reversic.Game)
	gr.POST("/:id/game.stat", reversic.StatGame)
	gr.POST("/:id/game.join", reversic.JoinGame)
	gr.GET("/:id/game.sub", reversic.SubGame)
	gr.POST("/:id/game.play/:name", reversic.PlayGame)

	app.Run(":2222")
}
