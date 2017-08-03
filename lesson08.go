package main

import (
	"fmt"
)

func main() {

	var veyron Vehicle
	veyron.id = 1
	veyron.name = "bugatti veyron"
	veyron.color = "black"

	var gayyardo Vehicle = Vehicle{2, "lomborghini gallardo", "gold"}
	sesto := Vehicle{id: 3, name: "ferrari testo elemento", color: "red"}

	write_vehicle_to_console(&veyron)
	write_vehicle_to_console(&gayyardo)
	write_vehicle_to_console(&sesto)

	var someVehicles []Vehicle
	someVehicles = append(someVehicles, sesto)
	someVehicles = append(someVehicles, gayyardo)
	someVehicles = append(someVehicles, veyron)

	write_all_vehicle_to_console(someVehicles)

	veyron.move(12, 10, -8)

	for _, vehicle := range someVehicles {
		vehicle.move(-3, 4, 10) // some_vehicles slice'ı içerisindeki tüm Vehicle nesneleri üzerinden move metodunu çağırıyoruz
	}

	var motto word = "it's a lovely day"
	fmt.Printf(motto.write_with_space())
}

type word string

type Vehicle struct {
	id          int
	name, color string
}

func write_vehicle_to_console(vehicle *Vehicle) {
	fmt.Printf("Id: %d name: %s color: %s\n", vehicle.id, vehicle.name, vehicle.color)
}

func write_all_vehicle_to_console(vehicles []Vehicle) {
	for _, vehicle := range vehicles {
		write_vehicle_to_console(&vehicle)
	}
}

func (vehicle Vehicle) move(x, y, z int) {
	fmt.Printf("%s, (%d:%d:%d) lokasyonuna gidiyor\n", vehicle.name, x, y, z)
}

func (word word) write_with_space() string {
	var result string = "aa"
	for _, c := range word {
		result += string(c) + " "
	}
	return result
}
