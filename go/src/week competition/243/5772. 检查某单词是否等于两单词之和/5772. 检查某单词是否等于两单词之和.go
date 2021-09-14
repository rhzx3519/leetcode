package main

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return alphabeta(firstWord) + alphabeta(secondWord) == alphabeta(targetWord)
}

func alphabeta(s string) int32 {
	var i int32
	for _, c := range s {
		i = i*10 + (c - 'a')
	}
	return i
}