package main

import "fmt"

func (W *World) useTool(tool string, subject string) {
	p := WORLD.Person
	if tool == "ключи" {
		lock := subject
		if lock == "дверь" {
			if p.backPackUp {
				if p.backPack.inStore("ключи") {
					keys, _ := getAllThing(p.backPack.things, "ключи")
					// key, _ := getThing(p.backPack.things, "ключи")
					lockedDooors := (p.Room).(*Room).getAllLockedDoors()
					if lockedDooors == nil {
						fmt.Println("все двери открыты")
						return
					}
					for _, door := range lockedDooors {
						for _, key := range keys {
							unlock := door.OpenLock(*(key.(*Key)))
							if unlock {
								fmt.Println("дверь открыта")
								break
							} else {
							}
						}
						fmt.Println("нет подходящих ключей")
					}

				} else {
					fmt.Println("нет предмета в инвентаре - ключи")
				}
			} else {
				fmt.Println("нет предмета в инвентаре - ключи")
			}
		} else if lock == "шкаф" {
			st := (p.Room.Store()).(*Storage)
			t, ok := getThing(st.things, lock)
			if ok {
				k := t.(*Container)
				key, ok := getThing(st.things, tool)
				if k.locked && ok {
					k.OpenLock(*(key.(*Key)))
					fmt.Println("Шкаф открыт")
					return
				} else {
					fmt.Println("Этот шкаф уже открыт или нет ключей")
				}
			} else {
				fmt.Println("не к чему применить")
			}
		}
	} else {
		fmt.Println("нет предмета в инвентаре - ", tool)
	}
}
