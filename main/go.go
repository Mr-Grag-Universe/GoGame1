package main

import "fmt"

func (W *World) goPlace(place string) {
	p := W.Person
	if canGo(p.Room, place) {
		room := roomGo(p.Room.Sib(), place)
		WORLD.Person.Move(room)
		report := room.enterIn()
		fmt.Println(report)
	} else {
		fmt.Println("нет пути в " + place)
	}
}
