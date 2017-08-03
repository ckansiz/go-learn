package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	gridZone := Game{
		5555,
		"Mohani Gezegeni Görevi",
		[]Player{
			Player{100, "deli", "cevat", 10.90},
			Player{102, "nadya", "komanaççi", 12.45},
			Player{103, "biricit", "bardot", 900.45},
		},
	}

	jsonOutput, _ := json.Marshal(gridZone) // jsonOutput json sonucunu , _ ise hata var ise onu döndürür
	fmt.Println(string(jsonOutput))

	var game Game
	if err := json.Unmarshal(jsonOutput, &game); err != nil {
		panic(err)
	}

	fmt.Printf("Game : %s\n", game.Name)
	for _, player := range game.Players {
		fmt.Println(player.FirstName, player.LastName, player.Point)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(game)

}

type Player struct {
	Id        int
	FirstName string
	LastName  string
	Point     float32
}

type Game struct {
	Id      int
	Name    string
	Players []Player
}
