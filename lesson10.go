package main

import (
	"fmt"
)

func main() {
	vosvos := Car{id: 10, name: "Orange vos vos"}
	cesna := Plane{id: 900, owner: "Chesna", maxAltitude: 3000}
	pejo205 := Car{id: 12, name: "Pejo 205 GTI"}

	startVehicleEngine(vosvos)
	startVehicleEngine(cesna)
	startVehicleEngine(pejo205)

	allTeam := []VehicleContract{
		vosvos,
		Car{id: 2, name: "Lilte Dori FSI"},
		cesna,
		pejo205,
		Plane{id: 905, owner: "Airbus", maxAltitude: 4000},
	}

	for _, member := range allTeam {
		member.startEngine()
	}
}

func (car Car) startEngine() {
	fmt.Printf("%s motoru çalıştırılıyor\n", car.name)
}

func (plane Plane) startEngine() {
	fmt.Printf("%s motoru çalıştırılıyor\n", plane.owner)
}

func startVehicleEngine(vehicle VehicleContract) {
	vehicle.startEngine()
}

type VehicleContract interface {
	startEngine()
}

type Car struct {
	id   int
	name string
}

type Plane struct {
	id, maxAltitude int
	owner           string
}
