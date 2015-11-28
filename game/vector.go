package game

//Vector is a 2d vector
type Vector struct {
	X, Y int
}

//InBoundsOf checks if this vector is in bounds of other vector
func (v Vector) InBoundsOf(other Vector) bool {
	return v.X < other.X && v.Y < other.Y
}

//Add adds two vectors together
func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}
