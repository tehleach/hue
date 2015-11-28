package game

import (
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/tehleach/hue/errors"
	"github.com/tehleach/hue/rest"
)

//Controller is the game controller
type Controller struct {
	rest.Controller
	DB *mgo.Database `inject:""`
}

//SetRoutes sets the controllers routes in given router
func (c *Controller) SetRoutes(router *httprouter.Router) {
	router.GET("/games", c.ListGames)
	router.GET("/games/:id", c.GetGame)
	router.POST("/games", c.NewGame)
}

//ListGames lists all games
func (c *Controller) ListGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	collection := c.DB.C("games")
	var games []Game
	if err := collection.Find(nil).All(&games); err != nil {
		c.Error(w, errors.New("something went wrong"))
		return
	}
	fmt.Fprint(w, games)
}

//NewGame intializes a game
func (c *Controller) NewGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	collection := c.DB.C("games")
	game := Game{Board: NewBoard(Vector{5, 5})}
	game.Board.PlacePiece(Vector{0, 0})
	game.Board.PlacePiece(Vector{4, 4})
	if err := collection.Insert(&game); err != nil {
		c.Error(w, errors.New("something went wrong"))
		return
	}
	fmt.Fprint(w, game)
}

//GetGame lists specific game
func (c *Controller) GetGame(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	collection := c.DB.C("games")
	var game Game
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		c.Error(w, errors.NewNotFound("Game", "ID", id))
		return
	}
	if err := collection.FindId(bson.ObjectIdHex(id)).One(&game); err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.Error(w, errors.NewNotFound("Game", "ID", id))
			return
		}
		c.Error(w, err)
		return
	}
	fmt.Fprint(w, game.Board.GetCurrentState())
}
