package game

//Space is a space on the board
type Space struct {
	HasPiece bool
	Piece    Piece
}

func getEmptySpaceGrid(dimensions Coords) [][]Space {
	spaces := make([][]Space, dimensions.Y)
	for i := range spaces {
		spaces[i] = make([]Space, dimensions.X)
		for j := range spaces[i] {
			spaces[i][j] = Space{}
		}
	}
	return spaces
}
