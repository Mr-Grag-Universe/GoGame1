package main

func (r *Room) lookup() string {
	/*if r.Name == "улица" {
		//fmt.Println()
		return "на улице весна. можно пройти - домой"
	}

	res := r.Store.printAll()
	if r.Store.isEmpty() && r.Name != "коридор" {
		res = "пустая комната." + " " + res
	} else {
		res = r.info + " " + res
	}*/

	res := r.info
	x := r.store.printAll()
	if x != "" && x != " " {
		res += " " + x
	}
	/*res += ". можно пройти -"
	unlocked := 0
	for _, val := range r.Sib {
		if val.locked == false {
			if val.r1.Name != r.Name {
				res += " " + val.r1.Name
			} else {
				res += " " + val.r2.Name
			}
			unlocked++
		}
		if _, ul := r.NoLaUD(); ul != unlocked {
			res += ","
		}
	}*/

	res = deleteSpaces(res)

	return res
	//r.Store.p
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) getAllLockedDoors() []*Door {
	var doors []*Door
	for _, val := range r.Sib() {
		if val.IsLocked() {
			doors = append(doors, val)
		}
	}
	return doors
}

func (r *Room) NoLaUD() (int, int) {
	numberOfUnlD := 0
	numberOfLD := 0
	for _, val := range r.Sib() {
		if val.IsLocked() {
			numberOfLD++
		} else {
			numberOfUnlD++
		}
	}
	return numberOfLD, numberOfUnlD
}

func (r *Room) enterIn() string {
	res := ""
	if r.store.isEmpty() == false || r.name == "коридор" || r.name == "улица" {
		res += r.info
	} else {
		res += "пустая комната. "
	}

	res += " можно пройти - "
	for _, val := range r.Sib() {
		if val.r1.Name() != r.name {
			res += val.r1.Name() + ", "
		} else {
			res += val.r2.Name() + ", "
		}
	}

	return res
}

func (street *Street) Name() string {
	return street.name
}

func (street *Street) Address() string {
	return street.address
}

func housesNames(houses []*House) string {
	res := ""
	for _, val := range houses {
		res += " " + val.name + ","
	}
	return res
}

func (street *Street) lookup() string {
	return street.info + "адрес улицы: " + street.address + ". можно пройти -" + housesNames(street.houses)
}

func (street *Street) enterIn() string {
	return street.info + " можно пройти -" + housesNames(street.houses)
}

func (street *Street) Sib() []*Door {
	var res []*Door
	for _, house := range street.houses {
		res = append(res, house.enter)
	}
	return res
}

func (street *Street) Store() Storages {
	return nil
}

func (r *Room) Sib() []*Door {
	return r.sib
}

func (r *Room) Store() Storages {
	return &r.store
}
