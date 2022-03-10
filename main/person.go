package main

type Person struct {
	Name       string
	Room       Rooms
	backPack   BackPack
	backPackUp bool
}

func (person *Person) Move(room Rooms) {
	person.Room = room
}

func lookup(r *Room) {

}
