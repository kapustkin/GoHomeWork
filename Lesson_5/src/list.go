package main	
import (
	//"fmt"
	//"errors"
	//"log"
)


//List список хранит элемент с указателем на первый и последний элемент и количество элементов
type List struct {
	rootItem Item
	length  int     
}

//Item элемент списка
type Item struct {
	next, prev *Item
	list *List
	value interface{}
}

// insert - вставляет элемент item после элемента next
func (list *List) insert (item, next *Item) {
	nextItem := next.next
	next.next = item
	item.prev = next
	item.next = nextItem
	nextItem.prev = item
	item.list = list
	list.length++
}

func (list *List) checkRootItem() {
	if list.rootItem.next == nil {
		list.rootItem.next = &list.rootItem
		list.rootItem.prev = &list.rootItem
	}
}


// PushFront - добавляет элемент в начало
func (list *List) PushFront(value interface{}) {
	list.checkRootItem()
	list.insert(&Item{value: value}, &list.rootItem)
}

// PushBack - добавляет элемент в конец
func (list *List) PushBack(value interface{}) {
	list.checkRootItem()
	list.insert(&Item{value: value}, list.rootItem.prev)
}

// First возвращает первый элемент
func (list *List) First()(*Item)  {
	return list.rootItem.next 
}

// Last возвращает последний элемент
func (list *List) Last()(*Item)  {
	return list.rootItem.prev
}

// Len - возвращает количество Item в списке
func (list *List) Len() int {
	if list == nil {
		return 0
	}
	return list.length
}

// Value - возвращает значение Item
func (item *Item) Value() (result interface{}) {
	if item != nil {
		result = item.value
	}
	return
}

// Next - возвращает следующий элемент
func (item *Item) Next() (result *Item) {
	if item != nil && item.list != nil && item.next != &item.list.rootItem {
		result = item.next
	}
	return
}

// Prev - возвращает предыдущий элемент
func (item *Item) Prev() (result *Item) {
	if item != nil && item.list != nil && item.prev != &item.list.rootItem {
		result = item.prev
	}
	return
}

// Remove - удаляет текущий элемент из списка
func (item *Item) Remove() (result bool) {
	if item != nil && item.list != nil && item.list.rootItem.prev != nil {
		item.prev.next = item.next
		item.next.prev = item.prev
		item.list.length--
		item = nil
		result = true
	}
	return
}
