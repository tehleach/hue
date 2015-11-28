//Package game contains all game state
package game

import (
	"bytes"

	"gopkg.in/mgo.v2/bson"
)

//Game represents game state
type Game struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Board Board         `bson:""`
}

//Board is a game board
type Board struct {
	Spaces [][]Space
}

//NewBoard gets a new board
func NewBoard(width, height int) Board {
	return Board{getSpaceGrid(width, height)}
}

//GetCurrentState prints the current board state
func GetCurrentState(spaces [][]Space) string {
	var buffer bytes.Buffer

	for i := range spaces {
		for j := range spaces[i] {
			curSpace := spaces[i][j]
			if curSpace.HasPiece {
				buffer.WriteString("1")
			} else {
				buffer.WriteString("0")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
