package main

import (
	"fmt"
)

func main() {
	var zidane FootballPlayer
	zidane.self = Player{id: 10, nickname: "Zinedine Zidane"}
	zidane.position = "Midfield"
	zidane.abilities = []Ability{
		Ability{name: "shoot", power: 92},
		Ability{name: "high pass", power: 84},
	}

	zidane.abilities[1].useAbility()
	zidane.self.saySomething("What can i do sometimes. This is football")
	zidane.abilities[0].useAbility()

	var tayson Boxer
	tayson.self = Player{id: 88, nickname: "Bulldog"}
	tayson.knockdownCount = 32
	tayson.abilities = []Ability{
		Ability{name: "Defence", power: 76},
	}

	tayson.self.saySomething("I will win this game!")
	tayson.abilities[0].useAbility()

}

func (ability *Ability) useAbility() {
	fmt.Printf("[%s] yeteneği kullanılıyor. Güç: %d\n", ability.name, ability.power)
}

func (player *Player) saySomething(message string) {
	fmt.Printf("%s says that %s\n", player.nickname, message)
}

type Player struct {
	id       int
	nickname string
}

type Ability struct {
	name  string
	power int
}

type FootballPlayer struct {
	position  string
	self      Player
	abilities []Ability
}

type Boxer struct {
	knockdownCount int
	self           Player
	abilities      []Ability
}
