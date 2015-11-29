package game

//Vector is a 2d vector
type Vector struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//InBoundsOf checks if this vector is in bounds of other vector
func (v Vector) InBoundsOf(other Vector) bool {
	return v.X < other.X && v.Y < other.Y && v.X >= 0 && v.Y >= 0
}

//Add adds two vectors together
func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}
