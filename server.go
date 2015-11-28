package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/facebookgo/inject"
	"github.com/julienschmidt/httprouter"
	"github.com/tehleach/hue/game"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB("hue")
	gameController := &game.Controller{}

	inject.Populate(db, gameController)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/games", gameController.ListGames)
	router.GET("/games/:id", gameController.GetGame)
	router.POST("/games", gameController.NewGame)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index is the main route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "LETS PLAY SOME GAMEZ!")
}
