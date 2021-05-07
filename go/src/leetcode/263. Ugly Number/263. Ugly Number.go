package _63__Ugly_Number

func isUgly(n int) bool {
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n %5 ==0 {
		n /= 5
	}
	return n==1
}