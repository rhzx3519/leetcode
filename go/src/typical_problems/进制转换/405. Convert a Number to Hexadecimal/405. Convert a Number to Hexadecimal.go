/**
@author ZhengHao Lou
Date    2021/10/02
*/
package main

import "fmt"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var t uint32
	//var sign = 1
	if num < 0 {
		//sign = -1
		t = uint32(^(-num)) + 1
	} else {
		t = uint32(num)
	}

	arr := []uint32{}
	for t != 0 {
		arr = append(arr, t % 16)
		t /= 16
	}
	bytes := []byte{}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= 10 {
			bytes = append(bytes, byte(arr[i] - 10 + 'a'))
		} else {
			bytes = append(bytes, byte(arr[i] + '0'))
		}
	}

	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	toHex(1)
	toHex(26)
	toHex(-1)
}