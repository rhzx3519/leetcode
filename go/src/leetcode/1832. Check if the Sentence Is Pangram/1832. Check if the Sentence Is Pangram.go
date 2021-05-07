package main

func checkIfPangram(sentence string) bool {
	alphabet := make([]int, 26)
	for _, c := range sentence {
		alphabet[c - 'a']++
	}
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] == 0 {
			return false
		}
	}
	return true
}
