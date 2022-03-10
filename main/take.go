package main

import "fmt"

func (W *World) take(name string) {
	p := W.Person
	thingName := name
	if p.Room.Store().inStore(thingName) {
		if p.backPackUp == true {
			st := (WORLD.Person.Room.Store()).(*Storage)
			thing, ok := getThing(st.things, thingName)
			if ok {
				WORLD.Person.backPack.append(thing)
				WORLD.Person.Room.Store().pop(thingName)
				fmt.Println("предмет добавлен в инвентарь: " + thingName)
			} else {
				fmt.Println("Не получилось взять предмет")
			}
		} else {
			fmt.Println("некуда класть")
			return
		}
	} else {
		fmt.Println("нет такого")
	}
}
