//Package game contains all game state
package game

import (
	"bytes"

	"github.com/tehleach/hue/errors"
)

//Board is a game board
type Board struct {
	dimensions Coords
	Spaces     [][]Space
}

//NewBoard gets a new board
func NewBoard(dimensions Coords) Board {
	return Board{dimensions, getEmptySpaceGrid(dimensions)}
}

//GetCurrentState prints the current board state
func (b *Board) GetCurrentState() string {
	var buffer bytes.Buffer

	for i := range b.Spaces {
		for j := range b.Spaces[i] {
			curSpace := b.Spaces[i][j]
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

//PlacePiece attempts to place a piece at coords
func (b *Board) PlacePiece(coords Coords) error {
	if coords.X >= b.dimensions.X || coords.Y >= b.dimensions.Y {
		return errors.NewOutOfBounds("Board")
	}
	space := &b.Spaces[coords.X][coords.Y]
	space.HasPiece = true
	space.Piece = Piece{10}
	return nil
}
