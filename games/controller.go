package games

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/bin16/gamebox/games/reversi"

	"github.com/bin16/gamebox/eventutil"
	"github.com/bin16/gamebox/storeutil"
	"github.com/bin16/gamebox/userc"
	"github.com/gin-gonic/gin"
)

func Reversi() *GameController {
	return &GameController{
		Namespace: "reversi",
		Core:      &reversi.Game{},
	}
}

type GameCore interface {
	ID() string
	New() interface{}
	Load([]byte) (interface{}, error)
	Play(string, []int) int
	Join(string) int
	Data(string) map[string]interface{}
	Save() []byte
	Dict(int) error
}

type GameController struct {
	Namespace string
	Core      GameCore
}

func (gc *GameController) Setup(g *gin.Engine) {
	g.POST(fmt.Sprintf("/%s.create", gc.Namespace), gc.Create)
	gr := g.Group("/" + gc.Namespace)
	gr.GET("/", func(c *gin.Context) {
		c.File(path.Join("static", gc.Namespace, "index.html"))
	})
	gr.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		if sid, _ := c.Cookie("_sid"); sid != "" {
			eventutil.UnSubscribe(id, sid)
		}
		c.Next()
	}, func(c *gin.Context) {
		c.File(path.Join("static", gc.Namespace, "game.html"))
	})
	gr.GET("/:id/game.data", gc.Data)
	gr.POST("/:id/game.stat", gc.Data)
	gr.POST("/:id/game.join", gc.Join)
	gr.POST("/:id/game.play/", gc.Play)
	gr.POST("/:id/game.play/:name", gc.Play)
}

func (gc *GameController) Create(c *gin.Context) {
	g := gc.Core.New().(GameCore)
	username := userc.Username(c)
	if err := g.Dict(g.Join(username)); err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"ok":     0,
			"errors": []string{err.Error()},
		})
		return
	}

	g.Join(username)
	storeutil.Set(g.ID(), g.Save())
	c.JSON(http.StatusCreated, gin.H{
		"ok":   1,
		"game": g.Data(username),
	})
}

func (gc *GameController) Data(c *gin.Context) {
	g := gc.Core
	username := userc.Username(c)
	id := c.Param("id")
	if !storeutil.Has(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"ok": 0,
		})
		return
	}
	data := storeutil.Get(id)
	if g1, err := g.Load(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": 0,
		})
		return
	} else {
		g = g1.(GameCore)
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   1,
		"data": g.Data(username),
	})
}

func (gc *GameController) Join(c *gin.Context) {
	g := gc.Core
	id := c.Param("id")
	username := userc.Username(c)
	if !storeutil.Has(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"ok": 0,
		})
		return
	}
	data := storeutil.Get(id)
	if g1, err := g.Load(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": 0,
		})
		return
	} else {
		g = g1.(GameCore)
	}

	if err := g.Dict(g.Join(username)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    0,
			"error": err.Error(),
		})
	}
	storeutil.Set(g.ID(), g.Save())
	eventutil.Post(id, "join")

	c.JSON(http.StatusOK, gin.H{"ok": 1})
}

func (gc *GameController) Events(c *gin.Context) {
	g := gc.Core
	username := userc.Username(c)
	if err := g.Dict(g.Join(username)); err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"ok":     0,
			"errors": []string{err.Error()},
		})
		return
	}

	g.Join(username)
	storeutil.Set(g.ID(), g.Save())
	c.JSON(http.StatusCreated, gin.H{
		"ok": 1,
	})
}

func (gc *GameController) Play(c *gin.Context) {
	g := gc.Core
	id := c.Param("id")
	username := userc.Username(c)
	if !storeutil.Has(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"ok": 0,
		})
		return
	}
	data := storeutil.Get(id)
	if g1, err := g.Load(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": 0,
		})
		return
	} else {
		g = g1.(GameCore)
	}

	name := c.Param("name")
	p := g.Data(name)["side"].(int)
	n := indexOf(name)
	if err := g.Dict(g.Play(username, []int{p, n})); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"ok":    0,
			"error": err.Error(),
		})
		return
	}
	storeutil.Set(g.ID(), g.Save())
	eventutil.Post(id, "play")

	c.JSON(http.StatusAccepted, gin.H{
		"ok": 1,
	})
}

// indexOf translate c4 into 2*8+3 = 19
func indexOf(name string) int {
	t := strings.ToLower(name)
	m := int(rune(t[0]) - 'a')
	n := int(rune(t[1]) - '1')
	return m*8 + n
}
