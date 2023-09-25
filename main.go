package main

import (
	"map/storage/mp"
)

type Storage interface {
	Len() int64
	Add(value int64) int64
	RemoveByIndex(id int64)
	RemoveByValue(value int64)
	RemoveAllByValue(value int64)
	GetByIndex(id int64) (value int64, ok bool)
	GetByValue(value int64) (id int64, ok bool)
	GetAllByValue(value int64) (ids []int64, ok bool)
	GetAll() (values []int64, ok bool)
	Clear()
	Print()
}

func main() {
	var m Storage
	m = mp.NewMap()
	m.Add(1)
	m.Add(2)
	m.Add(3)
	m.Add(3)
	m.Clear()
	m.Print()
}
