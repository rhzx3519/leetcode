/**
@author ZhengHao Lou
Date    2021/12/19
*/
package main

import "fmt"

func addSpaces(s string, spaces []int) string {
	bytes := []byte{}
	//var offset int
	for i := range s {
		if len(spaces) != 0 && i == spaces[0] {
			bytes = append(bytes, ' ')
			//offset++
			spaces = spaces[1:]
		}
		bytes = append(bytes, s[i])
	}
	return string(bytes)
}

func main() {
	fmt.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8,13,15}))
	fmt.Println(addSpaces("icodeinpython", []int{1,5,7,9}))
	fmt.Println(addSpaces("spacing", []int{0,1,2,3,4,5,6}))
}
