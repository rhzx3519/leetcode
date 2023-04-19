package main

import "strings"

func printBin(num float64) string {
	b := &strings.Builder{}
	b.WriteString("0.")
	for i := 0; i < 6; i++ {
		num *= 2
		if num < 0 {
			b.WriteByte('0')
		} else {
			b.WriteByte('1')
			num--
			if num == 0 {
				return b.String()
			}
		}
	}
	return "ERROR"
}
