package main	

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
    type pair struct{
		in string
		out string
	}
	test := []pair{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"a10bcd", "aaaaaaaaaabcd"},
		{"a4bc2d5a10", "aaaabccdddddaaaaaaaaaa"},
		{"45", ""},
		{"\\4d3", "4ddd"},
		{"\\45d3", "44444ddd"},
		{"\\4\\5d3", "45ddd"},
		{"a\\4bc3d", "a4bcccd"},
		{"a\\4bc\\\\3d5f", "a4bc\\3dddddf"},
		}
	for _, t := range test {
		str := decode(t.in)
		if t.out ==  str{
			fmt.Printf("%s = %s | Result %s\n", t.out, str, "OK")
		} else {
			fmt.Printf("%s <> %s | Result %s\n", t.out, str, "FAIL")
		}
	}
}

func decode(data string) (result string) {	
	var buffer string

	for pos, item := range data  {
		var ch = string(item)

		if pos == 0 { 
			if (isDigit(ch)) {
				return ""
			} 
			if ch != "\\" {
				result = ch
			}			
			continue
		}
		
		repChar := string(data[pos - len(buffer) - 1])
		if isDigit(ch) == true && repChar != "\\" {
			buffer = buffer + ch
			if (pos != len(data) - 1) {
				continue
			}
		} 
		if (buffer != "") {
			if (repChar == "\\") {
				result = result + buffer + ch;					
			} else {
				num, _ := strconv.ParseInt(buffer, 10, 64)
				result =  result + strings.Repeat(repChar, int(num -1))
				if !isDigit(ch) {
					result = result + ch
				} 
			}			
			buffer = ""		
		} else {
			if (ch != "\\") {
				result = result + ch
			} else {
				if (string(data[pos - len(buffer) - 1]) == "\\") {
					result = result  + "\\"
				}
			}
		}		
	}
	return
}

func isDigit(el string) bool {
	dig := []string{ "0","1","2","3","4","5","6","7","8","9" }
	for _, n := range dig {
		if el == n {
			return true
		}
	}
	return false
}