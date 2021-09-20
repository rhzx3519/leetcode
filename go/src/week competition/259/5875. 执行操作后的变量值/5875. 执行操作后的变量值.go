package main

func finalValueAfterOperations(operations []string) int {
	var x int
	for _, op := range operations {
		switch op {
		case "--X":
			x--
		case "X--":
			x--
		case "++X":
			x++
		case "X++":
			x++
		}
	}
	return x
}
