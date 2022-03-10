package main

func (door *Door) OpenLock(key Key) bool {
	for _, val := range key.doors {
		if door == val {
			door.locked = false
			return true
		}
	}
	return false
}
func (door *Door) IsLocked() bool {
	return door.locked
}

func (t *Thing) Name() string {
	return t.name
}

func (t *Key) Name() string {
	return t.name
}
