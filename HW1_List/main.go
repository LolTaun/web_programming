package main

import (
	"fmt"
	"list/storage/list"
	//"fmt"
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

	//ids, ok := l.GetAllByValue(11)
	//if ok {
	//	for _, val := range ids {
	//		print(val, " ")
	//	}
	//}
	fmt.Println(l.GetAll())
	l.Print()
}
