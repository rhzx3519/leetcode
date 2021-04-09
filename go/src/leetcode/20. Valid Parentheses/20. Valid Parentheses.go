package main

var mp = map[uint8]uint8 {
	')' : '(',
	'}' : '{',
	']' : '[',
}

func isValid(s string) bool {
	st := []uint8{}
	for i, _ := range s {
		if s[i]=='(' || s[i]=='{' || s[i]=='[' {
			st = append(st, s[i])
		} else {
			if len(st)==0 || st[len(st)-1] != mp[s[i]] {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}