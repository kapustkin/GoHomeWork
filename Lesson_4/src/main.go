package main	

import (
	"fmt"
	"errors"
	"log"
)

func main() {	
	data := []interface{}{1,2,3,4,8,3}

	maxElement, err := GetMaxItem(data, func(currentItemIndex, maxItemIndex int) bool {
		return data[currentItemIndex].(int) > data[maxItemIndex].(int)
	})
	
	if err != nil {
		log.Fatal(err)
	} else 
	{
		fmt.Printf("%+v\n", maxElement)
	}
}

//GetMaxItem Возвращает наибольший элемент из слайса на освнове компаратора
func GetMaxItem(data []interface{}, comparator func(currentItemIndex, maxItemIndex int) bool) (result interface{}, err error) {
	if len(data) == 0 {
		err = errors.New("Input slice is empty")
	} else {
		var maxItemIndex int
		for pos := range data {
			if comparator(pos, maxItemIndex) {
				maxItemIndex = pos
			}
		}
		result = data[maxItemIndex]
	}
	return
}