package main

var rand7 func() int

func rand10() int {
	// (0 ~ 6)*7+(0 ~ 6) = 0 ~ 48
	var t int
	for {
		t = (rand7() - 1) * 7 + rand7() - 1
		if t < 40 {
			return t%10 + 1
		}
	}

}
