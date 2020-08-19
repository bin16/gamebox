package reversic

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/bin16/Reversi/eventutil"
	"github.com/bin16/Reversi/reversi"
	"github.com/bin16/Reversi/storeutil"
	"github.com/bin16/Reversi/userc"
)

// Index GET /
func Index(c *gin.Context) {
	c.File("static/reversi/index.html")
}

// Game GET /:id
func Game(c *gin.Context) {
	id := c.Param("id")
	uid, _ := uuid.NewRandom()
	if sid, err := c.Cookie("_sid"); err != nil {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("_sid", uid.String(), 999999999, "/", "", false, true)
	} else {
		eventutil.UnSubscribe(id, sid)
	}

	c.File("static/reversi/game.html")
}

// CreateGame POST /game.create
func CreateGame(c *gin.Context) {
	username := userc.Username(c)
	g := reversi.NewGame()
	g.Join(username)
	storeutil.Set(g.ID, g.Save())
	c.JSON(http.StatusCreated, gin.H{
		"ok":   1,
		"game": g,
	})
}

// GameEvents GET /reversi/:id/game.events/ Server-Sent Events
func GameEvents(c *gin.Context) {
	id := c.Param("id")
	sid, _ := c.Cookie("_sid")
	sub := eventutil.Subscribe(id, sid)
	c.Stream(func(w io.Writer) bool {
		select {
		case msg, more := <-sub.OnMessage():
			if more && msg != nil {
				c.SSEvent("event", msg)
			}
			return more
		}
	})
}

// StatGame POST /:id/game.stat
func StatGame(c *gin.Context) {
	username := userc.Username(c)
	id := c.Param("id")
	if !storeutil.Has(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"ok": 0,
		})
		return
	}
	data := storeutil.Get(id)
	g, err := reversi.Load(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": 0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   1,
		"data": g.Data(username),
	})
}

// JoinGame POST /:id/game.join
func JoinGame(c *gin.Context) {
	id := c.Param("id")
	username := userc.Username(c)
	if !storeutil.Has(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"ok": 0,
		})
		return
	}
	data := storeutil.Get(id)
	g, err := reversi.Load(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": 0,
		})
		return
	}

	if err := g.Join(username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    0,
			"error": err.Error(),
		})
	}
	storeutil.Set(g.ID, g.Save())
	eventutil.Post(id, "join")

	c.JSON(http.StatusOK, gin.H{"ok": 1})
}

// PlayGame POST /:id/game.play/:name
func PlayGame(c *gin.Context) {
	id := c.Param("id")
	username := userc.Username(c)
	if !storeutil.Has(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"ok": 0,
		})
		return
	}
	data := storeutil.Get(id)
	g, err := reversi.Load(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": 0,
		})
		return
	}

	name := c.Param("name")

	if err := g.Play(username, name); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"ok":    0,
			"error": err.Error(),
		})
		return
	}
	storeutil.Set(g.ID, g.Save())
	eventutil.Post(id, "play")

	c.JSON(http.StatusAccepted, gin.H{
		"ok": 1,
	})
}
