package game

//Move represents a single move of a piece on a board
type Move struct {
	PieceCoords Vector `json:"pieceCoords"`
	Vector      Vector `json:"vector"`
}
