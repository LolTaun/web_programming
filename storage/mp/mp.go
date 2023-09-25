package mp

import "fmt"

type Map struct {
	nextIndex int64
	mp        map[int64]int64
}

func NewMap() *Map {
	return &Map{mp: make(map[int64]int64)}
}

// Len возвращает длину списка
func (m *Map) Len() (l int64) {
	return int64(len(m.mp))
}

// Add добавляет элемент в список и возвращает его индекс
func (m *Map) Add(value int64) (id int64) {
	m.mp[m.nextIndex] = value
	m.nextIndex++
	return m.nextIndex - 1
}

// RemoveByIndex удаляет элемент из списка по индексу
func (m *Map) RemoveByIndex(id int64) {
	delete(m.mp, id)
}

// RemoveByValue удаляет элемент из списка по значению
func (m *Map) RemoveByValue(value int64) {
	for k, v := range m.mp {
		if v == value {
			delete(m.mp, k)
			return
		}
	}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (m *Map) RemoveAllByValue(value int64) {
	for k, v := range m.mp {
		if v == value {
			delete(m.mp, k)
		}
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (m *Map) GetByIndex(id int64) (value int64, ok bool) {
	value, ok = m.mp[id]
	return
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (m *Map) GetByValue(value int64) (id int64, ok bool) {
	for k, v := range m.mp {
		if v == value {
			return k, true
		}
	}
	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (m *Map) GetAllByValue(value int64) (ids []int64, ok bool) {
	ids = make([]int64, 0)
	for k, v := range m.mp {
		if v == value {
			ids = append(ids, k)
			return ids, true
		}
	}
	return nil, false
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (m *Map) GetAll() (values []int64, ok bool) {
	if len(values) == 0 {
		return nil, false
	}
	values = make([]int64, 0)
	for _, v := range m.mp {
		values = append(values, v) // ?
	}
	return values, true
}

// Clear очищает список
func (m *Map) Clear() {
	m.mp = make(map[int64]int64)
	m.nextIndex = 0
}

// Print выводит список в консоль
func (m *Map) Print() {
	if len(m.mp) == 0 {
		fmt.Printf("Map is empty!\n")
		return
	}
	for k, v := range m.mp {
		fmt.Printf("%d:\t%d\n", k, v)
	}
}
