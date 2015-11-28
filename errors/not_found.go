package errors

import "fmt"

//NotFound is a not found error
type NotFound struct {
	Resource string
	Field    string
	Value    string
}

//NewNotFound returns a new NotFound
func NewNotFound(resource, field, value string) *NotFound {
	return &NotFound{resource, field, value}
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("%v with %v [%v] not found.", e.Resource, e.Field, e.Value)
}
