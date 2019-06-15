package main	

import (
	"fmt"
	"errors"
	"log"
)

func main() {
	var test = []struct{
		data string
		idx int	
	}{
		{data: "bla", idx: 1},
		{data: "blo", idx: 6},
		{data: "blu", idx: 3},
		{data: "bli", idx: 8},
		{data: "bly", idx: 0},
	}	
	
	var data []interface{}
	for _,v := range test {
		data = append(data, v)
	}

	maxElement, err := GetMaxItem(data, func(currentItemIndex, maxItemIndex int) bool {
		return test[currentItemIndex].idx > test[maxItemIndex].idx
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