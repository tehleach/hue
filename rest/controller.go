package rest

import (
	"encoding/json"
	"net/http"

	"github.com/tehleach/hue/errors"
)

//Controller is the base controller class
type Controller struct {
}

//UnmarshalBody pulls the json object out of the request body
func (c *Controller) UnmarshalBody(r *http.Request, destination interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(destination); err != nil {
		return errors.NewJSONDecode(err.Error())
	}
	return nil
}

//Error writes error to w and matches code based on switch in errors.go
func (c *Controller) Error(w http.ResponseWriter, err error) {
	WriteError(w, MatchErrorToCode(err), err.Error())
}
