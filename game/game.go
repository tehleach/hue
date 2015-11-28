package game

import "gopkg.in/mgo.v2/bson"

//Game represents game state
type Game struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Board Board
}
