/**
@author ZhengHao Lou
Date    2022/05/29
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode.cn/problems/sender-with-largest-word-count/
*/
func largestWordCount(messages []string, senders []string) string {
	var mx int
	var ans string
	var counter = make(map[string]int)
	for i := range messages {
		counter[senders[i]] += len(strings.Split(messages[i], " "))

	}

	fmt.Println(counter["EMoUs"], counter["FnZd"])
	for s := range counter {
		if counter[s] >= mx {
			if counter[s] > mx {
				mx = counter[s]
				ans = s
			} else if s > ans {
				ans = s
			}
		}
	}
	return ans
}

func main() {
	largestWordCount([]string{"tP x M VC h lmD", "D X XF w V", "sh m Pgl", "pN pa", "C SL m G Pn v", "K z UL B W ee", "Yf yo n V U Za f np", "j J sk f qr e v t", "L Q cJ c J Z jp E", "Be a aO", "nI c Gb k Y C QS N", "Yi Bts", "gp No g s VR", "py A S sNf", "ZS H Bi De dj dsh", "ep MA KI Q Ou"},
		[]string{"OXlq", "IFGaW", "XQPeWJRszU", "Gb", "HArIr", "Gb", "FnZd", "FnZd", "HArIr", "OXlq", "IFGaW", "XQPeWJRszU", "EMoUs", "Gb", "EMoUs", "EMoUs"})
}
