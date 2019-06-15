package main	

import (
	"fmt"
	"errors"
	"log"
)

//List asdasd
type List struct {
	current Item
	length  int     
}

//Item asdasda
type Item struct {
	next, prev *Item
	list *List
	Value interface{}
}

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

}

