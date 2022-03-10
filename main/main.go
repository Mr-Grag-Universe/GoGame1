package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ACTIONS = []string{"идти", "осмотреться", "взять", "надеть", "применить"}
var THINGS = []string{"чай", "сыр", "ключи", "рюкзак", "дверь", "телефон", "конспекты"}
var ROOMS = []string{"комната", "кухня", "коридор", "улица"}
var CONTAINERS = []string{"шкаф", "стул", "стол"}

var GameLogic = map[string][]string{
	"идти":  ROOMS,
	"взять": THINGS,
}

type World struct {
	//Rooms  []*Room
	houses []*House
	Person Person
}

type Command struct {
	command string
	params  []string
}

func canGo(room Rooms, name string) bool {
	for _, val := range room.Sib() {
		if val.r1.Name() == name || val.r2.Name() == name {
			if val.locked {
				return false
			}
			return true
		}
	}
	return false
}

func canTake(things []Things, name string) bool {
	for _, val := range things {
		if val.Name() == name {
			return true
		}
	}
	return false
}

func roomGo(doors []*Door, name string) Rooms {
	for _, door := range doors {
		if door.r1.Name() == name {
			return door.r1
		} else if door.r2.Name() == name {
			return door.r2
		}
	}
	return nil
}

func (command Command) Execute() {
	p := WORLD.Person
	switch command.command {
	case "осмотреться":
		{
			answer := p.Room.lookup()
			if p.Room.Name() == "кухня" && p.backPackUp == false {
				answer += " надо собрать рюкзак и идти в универ. "
			}
			fmt.Println(answer)
		}
	case "идти":
		{
			place := command.params[0]
			WORLD.goPlace(place)
			return
		}
	case "взять":
		{
			thingName := command.params[0]
			WORLD.take(thingName)
			return
		}
	case "надеть":
		{
			thingName := command.params[0]
			WORLD.putOn(thingName)
		}
	case "применить":
		{
			tool := command.params[0]
			subject := command.params[1]
			WORLD.useTool(tool, subject)
		}
	}
}

type BackPack struct {
	things []Things
}
type BackPacks interface {
	append(t Things)
	pop(name string)
	inStore(name string) bool
}

func (b *BackPack) append(t Things) {
	b.things = append(b.things, t)
}

func (b *BackPack) pop(name string) {
	b.things = RemoveFromThingsArr(b.things, name, 1)
}

func (b *BackPack) inStore(name string) bool {
	for _, val := range b.things {
		if val.Name() == name {
			return true
		}
	}
	return false
}

func (b *BackPack) Name() string {
	return "рюкзак"
}

var WORLD World

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Intro()
	initGame()

	for scanner.Scan() {
		line := scanner.Text()

		answer := handleCommand(line)
		// fmt.Println(answer)
		if answer == "goodbye" {
			break
		} else if answer == "not implemented" {
			continue
		}
	}
}

func Intro() {
	fmt.Println("приветствую Вас в своей игре.")
	fmt.Println("здесь вы можете исследовать мир, пользоваться предметами и многое другое")
	fmt.Println("чтобы сделать что-то используйте команды")
	fmt.Println("формат команд: действие параметр1 параметр2")
	fmt.Println("допустимые команды:")
	fmt.Printf("идти $место:\n\tкомната/кухня/улица")
	fmt.Println("взять $предмет")
	fmt.Println("применить $инструмент $объект")
	fmt.Println("надеть $предмет")
}

func checkWords(words []string) bool {
	DICT := append(append(append(THINGS, ACTIONS...), ROOMS...), CONTAINERS...)
	for _, val := range words {
		if StrInArr(DICT, val) == false {
			return false
		}
	}
	return true
}

func checkFormat(words []string) bool {
	c := len(words) //ArrLen(words)
	switch words[0] {
	case "идти":
		{
			c := Count(ROOMS, words[1])
			if c != 0 {
				return true
			}
			return false
		}
	case "применить":
		{
			if len(words) != 3 {
				return false
			}
			return true
		}
	case "надеть":
		{
			if c < 2 {
				return false
			}
			return true
		}
	case "осмотреться":
		{
			if c > 1 {
				return false
			}
			return true
		}
	case "взять":
		{
			if c < 2 {
				return false
			}
			return true
		}
	default:
		return false
	}
	return false
}

func checkCommand(command string) bool {
	words := strings.Split(command, " ")
	if checkWords(words) == false || checkFormat(words) == false {
		return false
	}

	return true
}

func handleCommand(line string) string {
	line = strings.ToLower(line)
	if line == "finish" {
		return "goodbye"
	}

	if checkCommand(line) == false { // поменяй потом
		fmt.Println("Wrong command!")
		return "not implemented"
	}

	words := strings.Split(line, " ")
	command := Command{words[0], words[1:]}

	command.Execute()

	return "all is fine"
}
