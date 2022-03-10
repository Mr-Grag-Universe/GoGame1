package main

import "fmt"

func (W *World) putOn(thingName string) {
	switch thingName {
	case "рюкзак":
		{
			W.Person.backPackUp = true
			W.Person.Room.Store().pop(thingName)
			fmt.Println("вы надели: рюкзак")
			return
		}
	default:
		{
			fmt.Println("Вы не можете это надеть")
		}
	}
}
