//Package game contains all game state
package game

import "bytes"

//Board is a game board
type Board struct {
	spaces [][]*Space
}

//NewBoard gets a new board
func NewBoard(width, height int) *Board {
	return &Board{getSpaceGrid(width, height)}
}

//GetCurrentState prints the current board state
func (b *Board) GetCurrentState() string {
	var buffer bytes.Buffer

	for i := range b.spaces {
		for j := range b.spaces[i] {
			curPiece := b.spaces[i][j]
			if curPiece != nil {
				buffer.WriteString("1")
			} else {
				buffer.WriteString("0")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
