package errors

import "fmt"

//OutOfBounds is an error for something that went out of bounds
type OutOfBounds struct {
	Resource string
}

//NewOutOfBounds returns a new OutOfBounds error
func NewOutOfBounds(resource string) *OutOfBounds {
	return &OutOfBounds{resource}
}

func (e *OutOfBounds) Error() string {
	return fmt.Sprintf("Tried to index %v out of its bounds.", e.Resource)
}
