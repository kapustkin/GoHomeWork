package main	

import (
	"fmt"
	"log"
)

const debug = false

func main() {
	//Проверка List
	assertLen()
	assertPushFront()
	assertPushBack()	
	assertFirst()
	assertLast()
	//Проверка Item
	assertNext()
	assertPrev()
	assertValue()
	assertRemove()
}

func assert(current, need interface{}) {
	if current == need {
		if (debug) {
			fmt.Printf("\n%s\t actual:%+v\texcepted:%+v", "PASSED", current, need)
		}
	} else {
		log.Fatalf("\n%s\t actual:%+v\texcepted:%+v", "FAILED", current, need)
	}
}

func assertLen() {
	var list List
	fmt.Printf("%s", "Проверка Len...")
	assert(list.Len(), 0)
	list.PushFront(2)
	assert(list.Len(), 1)
	list.PushFront(2)
	list.PushFront(2)
	assert(list.Len(), 3)
	fmt.Printf("%s\n", " OK")
}

func assertPushFront() {
	var list List
	fmt.Printf("%s", "Проверка PushFront...")
	list.PushFront(2)
	list.PushFront(1)
	assert(list.First().Value(), 1)
	list.PushFront("стр")
	assert(list.First().Value(), "стр")
	assert(list.First().Next().Next().Value(), 2)
	fmt.Printf("%s\n", " OK")
}

func assertPushBack() {
	var list List
	fmt.Printf("%s", "Проверка PushBack...")
	list.PushBack(2)
	list.PushBack(1)
	list.PushBack("стр")

	assert(list.First().Value(), 2)
	assert(list.First().Next().Value(), 1)
	assert(list.First().Next().Next().Value(), "стр")
	fmt.Printf("%s\n", " OK")
}

func assertFirst() {
	var list List
	fmt.Printf("%s", "Проверка First...")
	var item *Item
	assert(list.First(), item)
	list.PushBack(2)
	list.PushBack(1)
	assert(list.First().Value(), 2)
	list.PushFront(333)
	assert(list.First().Value(), 333)
	fmt.Printf("%s\n", " OK")
}

func assertLast() {
	var list List
	fmt.Printf("%s", "Проверка Last...")
	var item *Item
	assert(list.Last(), item)
	list.PushBack(2)
	list.PushBack(1)
	assert(list.Last().Value(), 1)
	list.PushFront(333)
	list.PushBack(4)
	assert(list.Last().Value(), 4)
	fmt.Printf("%s\n", " OK")
}

func assertNext() {
	var list List
	fmt.Printf("%s", "Проверка Next...")
	list.PushBack(2)
	list.PushBack(1)	
	list.PushBack(4)
	assert(list.First().Value(), 2)
	assert(list.First().Next().Value(), 1)
	assert(list.First().Next().Next().Value(), 4)
	var item *Item
	assert(item.Next().Value(), nil)
	fmt.Printf("%s\n", " OK")
}

func assertPrev() {
	var list List
	fmt.Printf("%s", "Проверка Prev...")
	list.PushBack(2)
	list.PushBack(1)	
	list.PushBack(4)
	assert(list.Last().Value(), 4)
	assert(list.Last().Prev().Value(), 1)
	assert(list.Last().Prev().Prev().Value(), 2)
	var item *Item
	assert(item.Prev().Value(), nil)
	fmt.Printf("%s\n", " OK")
}


func assertValue() {
	var list List
	fmt.Printf("%s", "Проверка Value...")
	list.PushBack(2)
	assert(list.Last().Value(), 2)
	list.PushBack("asddsasd")
	assert(list.Last().Value(), "asddsasd")
	var item *Item
	assert(item.Value(), nil)
	fmt.Printf("%s\n", " OK")
}

func assertRemove() {
	var list List
	fmt.Printf("%s", "Проверка Remove...")
	list.PushBack(2)
	assert(list.Len(), 1)
	list.First().Remove();
	assert(list.Len(), 0)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.First().Next().Remove();
	assert(list.First().Value(), 1)
	assert(list.First().Next().Value(), 3)
	assert(list.Last().Remove(), true)
	assert(list.First().Value(), 1)
	var item *Item
	assert(item.Remove(), false)
	
	fmt.Printf("%s\n", " OK")
}