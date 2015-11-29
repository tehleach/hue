package game

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//Game represents game state
type Game struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Board Board
}

//Info is info required to create a game
type Info struct {
	Players         []string `json:"players"`
	BoardDimensions Vector   `json:"boardDimensions"`
}

//PieceState is state about a piece
type PieceState struct {
	Coords Vector `json:"coords"`
	Piece  Piece  `json:"piece"`
}

//State is state about a game
type State struct {
	Pieces []Piece `json:"pieces"`
}

//PrettyPrint prints a readable game
func (g *Game) PrettyPrint() string {
	return fmt.Sprintf("ID: %v\n%v", g.ID.String(), g.Board.GetCurrentState())
}
