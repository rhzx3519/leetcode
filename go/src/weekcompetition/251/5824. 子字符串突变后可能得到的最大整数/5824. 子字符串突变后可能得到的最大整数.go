package main

import (
	"fmt"
)

func maximumNumber(num string, change []int) string {
	bytes := []byte{}
	var j = -1
	var f = false
	for i := range num {
		d := int(num[i] - '0')
		if change[d] > d {
			f = true
			bytes = append(bytes, byte('0' + change[d]))
		} else if change[d] < d {
			if f {
				j = i
				break
			} else {
				bytes = append(bytes, num[i])
			}
		} else {
			bytes = append(bytes, num[i])
		}
	}
	if j != -1 {
		bytes = append(bytes, num[j:]...)
	}
	return string(bytes)
}

func main() {
	fmt.Println(maximumNumber("132", []int{9,8,5,0,3,6,4,2,6,8}))
	fmt.Println(maximumNumber("021", []int{9,4,3,5,7,2,1,9,0,6}))
	fmt.Println(maximumNumber("5", []int{1,4,7,5,3,2,5,6,9,4}))
	fmt.Println(maximumNumber("334111", []int{0,9,2,3,3,2,5,5,5,5})) // "334999"
	fmt.Println(maximumNumber("132", []int{9,8,5,0,3,6,4,2,6,8})) // "832"
}

//"132"
//[9,8,5,0,3,6,4,2,6,8]
// "832"
