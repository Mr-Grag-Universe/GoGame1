package main

type Things interface {
	Name() string
}

type Thing struct {
	name string
}

type Key struct {
	name  string
	doors []*Door
}

type Locks interface {
	IsLocked() bool
	OpenLock(key Key) bool
}

type Door struct {
	r1     Rooms
	r2     Rooms
	locked bool
}

type Storages interface {
	append(t Things)
	pop(name string)
	inStore(name string) bool
	printAll() string
}

type Storage struct {
	things []Things
}
type Container struct {
	name   string
	locked bool
	Storage
}

type Rooms interface {
	lookup() string
	Name() string
	enterIn() string
	Sib() []*Door
	Store() Storages
	// getAllLockedDoors() []*Door
	// NoLaUD() (int, int)
}

type Room struct {
	name  string
	info  string
	store Storage
	sib   []*Door
}

type Street struct {
	name    string
	address string
	info    string
	houses  []*House
}

var RoomsNums = map[string]int{
	"комната": 0,
	"кухня":   1,
}

type House struct {
	name    string
	address string
	rooms   []*Room
	enter   *Door
}
