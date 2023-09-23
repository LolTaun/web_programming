package main

import (
	"fmt"
	"list/storage/list"
)

func main() {
	l := list.List{}
	l.Add(11)
	l.Add(22)
	l.Add(33)
	l.Add(44)
	l.Add(55)
	l.Add(11)
	l.Add(11)
	

	fmt.Println(l.GetAll())
	l.Print()
}
