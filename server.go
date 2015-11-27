package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tehleach/hue/game"
)

var board *game.Board

func main() {
	board = game.NewBoard(5, 5)
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/games", ListGames)
	router.GET("/games/:id", GetGame)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index is the main route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "LETS PLAY SOME GAMEZ!")
}

//ListGames lists all games
func ListGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Would print all games here!")
}

//GetGame lists specific game
func GetGame(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "State for game %v here!", p.ByName("id"))
}
