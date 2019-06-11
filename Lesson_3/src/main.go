package main	

import (
	"fmt"
	"strings"
	"sort"
)

const selectingLimit = 10

func main() {
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. `
	pars := GetMostPopularWords(text)
	for _, item := range pars {
		fmt.Printf("%s - %d\n", item.Key, item.Value)
	}
}
// KeyVal ключ - значение
type KeyVal struct {
	Key   string
	Value int
}

// GetMostPopularWords - Функция для нахождения 10 самых часто встречающихся слов без учета словоформ
func GetMostPopularWords(text string) []KeyVal {
	parsedText := parseText(text)
	return sortMap(parsedText)
}

func parseText(text string)(map[string]int ) {
		result := map[string]int{}
		for _, word := range strings.Split(text, " ") {
			val, res := result[word]
				if res {
					result[word] = val + 1
				} else {
					result[word] = 1
				}
		}
		return result
}

func sortMap(source map[string]int) []KeyVal {
	var result []KeyVal
	for key, val := range source {
		result = append(result, KeyVal{key, val})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})
	if (len(result) > selectingLimit) {
		return result[:selectingLimit]
	} 
	return result	
}