package game

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

//Controller is the game controller
type Controller struct {
	DB *mgo.Database `inject:""`
}

//ListGames lists all games
func (c *Controller) ListGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	collection := c.DB.C("games")
	var games []Game
	if err := collection.Find(nil).All(&games); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}
	fmt.Fprint(w, games)
}

//NewGame intializes a game
func (c *Controller) NewGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	collection := c.DB.C("games")
	game := Game{Board: NewBoard(5, 5)}
	if err := collection.Insert(&game); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not object id hex"))
		return
	}
	if err := collection.FindId(bson.ObjectIdHex(id)).One(&game); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprint(w, GetCurrentState(game.Board.Spaces))
}
