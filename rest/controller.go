package rest

import "net/http"

//Controller is the base controller class
type Controller struct {
}

func (c *Controller) Error(w http.ResponseWriter, err error) {
	WriteError(w, MatchErrorToCode(err), err.Error())
}
