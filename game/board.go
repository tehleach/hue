//Package game contains all game state
package game

import (
	"bytes"

	"github.com/tehleach/hue/errors"
)

//Board is a game board
type Board struct {
	Dimensions Vector
	Spaces     [][]Space
}

//NewBoard gets a new board
func NewBoard(dimensions Vector) Board {
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

//PlacePiece attempts to place a piece at vector
func (b *Board) PlacePiece(vector Vector) error {
	if vector.X >= b.Dimensions.X || vector.Y >= b.Dimensions.Y {
		return errors.NewOutOfBounds("Board")
	}
	space := &b.Spaces[vector.X][vector.Y]
	space.HasPiece = true
	space.Piece = Piece{10}
	return nil
}

//ApplyMove moves the piece according to move given
func (b *Board) ApplyMove(move Move) error {
	if !move.PieceCoords.InBoundsOf(b.Dimensions) {
		return errors.NewOutOfBounds("Board")
	}
	space := &b.Spaces[move.PieceCoords.X][move.PieceCoords.Y]
	if !space.HasPiece {
		return errors.New("No piece at space provided")
	}
	piece := space.Piece
	newLocation := move.PieceCoords.Add(move.Vector)
	if !newLocation.InBoundsOf(b.Dimensions) {
		return errors.NewOutOfBounds("Board")
	}
	newSpace := &b.Spaces[newLocation.X][newLocation.Y]
	newSpace.Piece = piece
	newSpace.HasPiece = true
	space.Piece = Piece{}
	space.HasPiece = false
	return nil
}
