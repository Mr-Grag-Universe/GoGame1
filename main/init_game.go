package main

func initGame() {
	cupboard := Container{name: "шкаф", locked: false}
	cupboard.Storage = Storage{
		things: []Things{&Thing{name: "чай"}},
	}

	chair := Container{name: "стул", locked: false}
	chair.Storage = Storage{
		things: []Things{&BackPack{things: nil}},
	}

	table := Container{name: "стол", locked: false}

	diningRoom := new(Room)
	myRoom := new(Room)
	hallway := new(Room)
	street := new(Street)

	door1 := new(Door)
	door2 := new(Door)
	door3 := new(Door)

	table.Storage = Storage{
		things: []Things{&Key{"ключи", []*Door{door1}}, &Thing{name: "конспекты"}},
	}

	*door1 = Door{hallway, street, true}
	*door2 = Door{r1: myRoom, r2: hallway, locked: false}
	*door3 = Door{r1: diningRoom, r2: hallway, locked: false}

	*diningRoom = Room{
		name: "кухня",
		info: "ты находишься на кухне,",
		sib:  []*Door{door3},
	}

	*myRoom = Room{
		name: "комната",
		info: "ты в своей комнате.",
		sib:  []*Door{door2},
	}

	*hallway = Room{
		name: "коридор",
		info: "ничего интересного.",
		sib:  []*Door{door1, door2, door3},
	}

	house := new(House)
	*house = House{
		name:    "дом",
		address: "",
		rooms:   []*Room{myRoom, diningRoom, hallway},
		enter:   nil,
	}

	/*
		*street = Room{
			name: "улица",
			info: "на улице весна.",
			Sib:  []*Door{door1},
		}*/
	*street = Street{
		name:   "улица",
		info:   "на улице весна.",
		houses: []*House{house},
	}

	myRoom.Store().append(&chair)
	myRoom.Store().append(&table)

	p := Person{Name: "Gay", Room: diningRoom, backPackUp: false}

	WORLD.houses = append(WORLD.houses, house)
	WORLD.Person = p

	handleCommand("осмотреться")
	handleCommand("завтракать")
	handleCommand("идти комната")
	handleCommand("идти коридор")
	handleCommand("применить ключи дверь")
	handleCommand("идти комната")
	handleCommand("осмотреться")
	handleCommand("взять ключи")
	handleCommand("надеть рюкзак")
	handleCommand("осмотреться")
	handleCommand("взять ключи")
	handleCommand("взять телефон")
	handleCommand("взять ключи")
	handleCommand("осмотреться")
	handleCommand("взять конспекты")
	handleCommand("осмотреться")
	handleCommand("идти коридор")
	handleCommand("идти кухня")
	handleCommand("осмотреться")
	handleCommand("идти коридор")
	handleCommand("идти улица")
	handleCommand("применить ключи дверь")
	handleCommand("применить телефон шкаф")
	handleCommand("применить ключи шкаф")
	handleCommand("идти улица")
}
