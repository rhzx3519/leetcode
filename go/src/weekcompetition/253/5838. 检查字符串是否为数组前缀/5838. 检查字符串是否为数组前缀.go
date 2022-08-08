package main

func isPrefixString(s string, words []string) bool {
	n := len(s)
	bytes := []byte{}
	for _, w := range words {
		bytes = append(bytes, []byte(w)...)
		if len(bytes) >= n {
			break
		}
		t := len(bytes)

		if string(bytes) != s[:t] {
			break
		}
	}

	return string(bytes) == s
}
