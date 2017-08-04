package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	port int = 8045
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("pages")))
	http.HandleFunc("/players", getPlayers)
	log.Println("Web Server is Active at Port: ", port)
	portString := ":" + strconv.Itoa(port)
	http.ListenAndServe(portString, nil)
}

func getPlayers(response http.ResponseWriter, request *http.Request) {
	log.Println("Get Request for Players")
	players := loadPlayers()
	log.Println("Players Loaded")
	t, err := template.ParseFiles("pages/players.html")
	if err == nil {
		t.Execute(response, players)
	} else {
		log.Println(err.Error())
	}
}

type Player struct {
	Id       int
	Nickname string
	Level    int
}

func loadPlayers() []Player {
	return []Player{
		Player{1001, "Molfiryin", 2},
		Player{1002, "Gul'dan", 21},
		Player{1003, "Anduin", 12},
		Player{1004, "Lexar", 5},
		Player{1005, "Turol", 34},
	}
}
