package main

func largestOddNumber(num string) string {
	i := len(num)-1
	for ; i >= 0; i-- {
		if (num[i] - '0') & 1 == 1 {
			break
		}
	}
	if i == -1 {
		return ""
	}
	return num[:i+1]
}
