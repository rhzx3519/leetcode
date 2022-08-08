/**
@author ZhengHao Lou
Date    2021/12/26
*/
package _963__反转两次的数字

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	return num % 10	!= 0
}
