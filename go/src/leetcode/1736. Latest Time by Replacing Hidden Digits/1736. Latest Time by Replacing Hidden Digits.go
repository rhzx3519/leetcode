package main

func maximumTime(time string) string {
	bytes := make([]byte, 0)

	for i := range time {
		c := time[i]
		t := c
		switch c {
		case '?':
			if i == 0 {
				if time[1] <= '3' || time[1] == '?' {
					t = '2'
				} else {
					t = '1'
				}
			} else if i == 1 {
				if bytes[0] == '2' {
					t = '3'
				} else {
					t = '9'
				}
			} else if i == 3 {
				t = '5'
			} else {
				t = '9'
			}
		default:
			break
		}
		bytes = append(bytes, t)
	}

	return string(bytes)
}
