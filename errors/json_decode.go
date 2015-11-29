package errors

import "fmt"

//JSONDecode is a json decoding error
type JSONDecode struct {
	Message string
}

//NewJSONDecode returns a new JSONDecode
func NewJSONDecode(message string) *JSONDecode {
	return &JSONDecode{message}
}

func (e *JSONDecode) Error() string {
	return fmt.Sprintf("Error decoding body JSON: %v.", e.Message)
}
