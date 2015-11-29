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
	router.POST("/games", c.NewGame)
	router.GET("/games/:id", c.PrintGameState)
	router.POST("/games/:id/move", c.ApplyMove)
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

//PrintGameState lists specific game
func (c *Controller) PrintGameState(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	game, err := c.findGame(id)
	if err != nil {
		c.Error(w, err)
		return
	}
	fmt.Fprint(w, game.Board.GetCurrentState())
}

//ApplyMove applies a move to specific game
func (c *Controller) ApplyMove(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var move Move
	if err := c.UnmarshalBody(r, &move); err != nil {
		c.Error(w, err)
		return
	}

	id := p.ByName("id")
	game, err := c.findGame(id)
	if err != nil {
		c.Error(w, err)
		return
	}

	if err := game.Board.ApplyMove(move); err != nil {
		c.Error(w, err)
		return
	}
	if err := c.DB.C("games").UpdateId(bson.ObjectIdHex(id), game); err != nil {
		c.Error(w, err)
		return
	}
	fmt.Fprint(w, game.Board.GetCurrentState())
}

func (c *Controller) findGame(id string) (*Game, error) {
	collection := c.DB.C("games")
	var game Game
	if !bson.IsObjectIdHex(id) {
		return nil, errors.NewNotFound("Game", "ID", id)
	}
	if err := collection.FindId(bson.ObjectIdHex(id)).One(&game); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.NewNotFound("Game", "ID", id)
		}
		return nil, err
	}
	return &game, nil
}
