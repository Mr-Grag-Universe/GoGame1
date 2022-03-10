package main

import "strings"

func ArrLen(a ...interface{}) int {
	count := 0
	for _, _ = range a {
		count++
	}
	return count
}

func Count(arr []string, x string) int {
	i := 0
	for _, val := range arr {
		if val == x {
			i++
		}
	}
	return i
}

func StrInArr(arr []string, x string) bool {
	for _, val := range arr {
		if val == x {
			return true
		}
	}
	return false
}

func RemoveFromStrArr(arr []string, str string, n int) []string {
	copy := arr
	var res []string
	for _, val := range copy {
		if n > 0 {
			if val == str {
				n--
				continue
			} else {
				res = append(res, val)
			}
		}
	}
	return res
}

func RemoveFromThingsArr(arr []Things, name string, n int) []Things {
	copy := arr
	var res []Things
	for _, val := range copy {
		if n > 0 {
			if val.Name() == name {
				n--
				continue
			} else {
				p, ok := (val).(*Container)
				if ok && (*p).locked == false {
					isInStore := (*p).inStore(name)
					if isInStore {
						(*p).things = RemoveFromThingsArr((*p).things, name, n)
					}
					tmp := Things(p)
					res = append(res, tmp)
				} else {
					res = append(res, val)
				}
			}
		} else {
			res = append(res, val)
		}
	}
	return res
}

func getThing(things []Things, name string) (Things, bool) {
	for _, val := range things {
		if val.Name() == name {
			return val, true
		}
		p, ok := (val).(*Container)
		if ok && (*p).locked == false {
			x, ok := getThing((*p).things, name)
			if ok {
				return x, true
			}
		}
	}
	return nil, false
}

func getAllThing(things []Things, name string) ([]Things, bool) {
	var res []Things
	for _, val := range things {
		if val.Name() == name {
			res = append(res, val)
		}
		p, ok := (val).(*Container)
		if ok && (*p).locked == false {
			x, ok := getThing((*p).things, name)
			if ok {
				res = append(res, x)
			}
		}
	}
	return res, false
}

func deleteSpaces(line string) string {
	words := strings.Split(line, " ")
	//words = RemoveFromStrArr(words, "", len(words))
	res := ""
	for i, val := range words {
		res += val
		if val != words[i] {
			if words[i+1] != "." {
				res += " "
			}
		} else {
			res += " "
		}
	}
	return res
}
