package game

//Space is a space on the board
type Space struct {
	Piece *Piece
}

func getSpaceGrid(width, height int) [][]*Space {
	spaces := make([][]*Space, height)
	for i := range spaces {
		spaces[i] = make([]*Space, width)
	}
	return spaces
}

//Piece is a playable piece on the board
type Piece struct {
	HP int
}
