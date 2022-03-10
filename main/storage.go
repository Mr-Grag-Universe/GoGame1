package main

func (s *Container) Name() string {
	return s.name
}
func (s *Container) IsLocked() bool {
	return s.locked
}
func (s *Container) OpenLock(key Key) bool {
	s.locked = false
	return true
}

func (s *Storage) append(t Things) {
	s.things = append(s.things, t)
	//panic("implement me")
}

func (s *Storage) pop(name string) {
	s.things = RemoveFromThingsArr(s.things, name, 1)
	//panic("implement me")
}

func (s *Storage) inStore(name string) bool {
	for _, val := range s.things {
		if val.Name() == name {
			return true
		}
		p, ok := (val).(*Container)
		if ok && (*p).locked == false {
			x := (*p).inStore(name)
			if x {
				return true
			}
		}
	}
	return false
}

func (s *Storage) printAll() string {
	res := ""
	for _, val := range s.things {
		//fmt.Printf("%s ", val.Name())
		switch val.(type) {
		case *Container:
			{
				if (val).(*Container).isEmpty() == false {
					res += val.Name() + ": "
				}
				res += ((val).(*Container)).Storage.printAll()
			}
		default:
			{
				res += val.Name()
			}
		}
		res += " "
	}
	return res
}

func (s *Storage) isEmpty() bool {
	for _, val := range s.things {
		switch val.(type) {
		case *Container:
			{
				x := (val).(*Container).isEmpty()
				if x == false {
					return false
				}
			}
		default:
			{
				return false
			}
		}
	}
	return true
}
