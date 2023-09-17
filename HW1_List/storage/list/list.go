package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() (l *List) {
	return &List{len: 0, firstNode: nil}
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	newNode := &node{value: value}

	if l.firstNode == nil {
		newNode.index = 0
		l.firstNode = newNode
		l.len++
		return 0
	}

	l.len++
	currentNode := l.firstNode
	for ; currentNode.next != nil; currentNode = currentNode.next {
	}
	newNode.index = currentNode.index + 1
	currentNode.next = newNode

	return currentNode.index + 1
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if l.firstNode == nil {
		return
	}
	if l.firstNode.index == id {
		l.firstNode = l.firstNode.next
		l.len--
		return
	}
	prevNode := l.firstNode
	for currentNode := l.firstNode.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.index == id {
			prevNode.next = currentNode.next
			l.len--
			return
		}
		prevNode = currentNode
	}
	return
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	if l.firstNode == nil {
		return
	}

	if l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
		return
	}
	prevNode := l.firstNode
	for currentNode := l.firstNode.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			prevNode.next = currentNode.next
			l.len--
			return
		}
		prevNode = currentNode
	}
	return
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	if l.firstNode == nil {
		return
	}

	if l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
	}
	prevNode := l.firstNode
	for currentNode := l.firstNode.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			prevNode.next = currentNode.next
			l.len--
			currentNode = prevNode
		}
		prevNode = currentNode
	}
	return
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if l.firstNode == nil {
		return 0, false
	}
	if l.firstNode.index == id {
		return l.firstNode.value, true
	}
	for currentNode := l.firstNode.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.index == id {
			return currentNode.value, true
		}
	}
	return 0, false
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	if l.firstNode == nil {
		return 0, false
	}

	if l.firstNode.value == value {
		return l.firstNode.index, true
	}
	for currentNode := l.firstNode.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			return currentNode.index, true
		}
	}
	return
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	if l.firstNode == nil {
		return nil, false
	}

	if l.firstNode.value == value {
		ids = append(ids, l.firstNode.index)
	}
	for currentNode := l.firstNode.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			ids = append(ids, currentNode.index)
		}
	}
	if len(ids) == 0 {
		return nil, false
	}
	return ids, true
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.firstNode == nil {
		return nil, false
	}
	values = append(values, l.firstNode.value)

	for currentNode := l.firstNode.next; currentNode != nil; currentNode = currentNode.next {
		values = append(values, currentNode.value)
	}
	return values, true
}

// Clear очищает список
func (l *List) Clear() {
	l.firstNode = nil // Any memory leaks?
	return
}

// Print выводит список в консоль
func (l *List) Print() {
	if l.firstNode == nil {
		fmt.Println("Empty")
		return
	}
	n := l.firstNode
	fmt.Print("[")
	for ; n != nil; n = n.next {
		fmt.Print(n.value)
		fmt.Print(", ")
	}
	fmt.Println("\b\b]")
}
