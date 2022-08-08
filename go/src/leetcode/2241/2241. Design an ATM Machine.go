/**
@author ZhengHao Lou
Date    2022/04/18
*/
package main

/**
https://leetcode-cn.com/problems/design-an-atm-machine/
*/

var values = []int{20, 50, 100, 200, 500}

type ATM struct {
	cashes []int
}

func Constructor() ATM {
	caches := make([]int, len(values))
	return ATM{
		cashes: caches,
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := range banknotesCount {
		this.cashes[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	var ans = make([]int, len(values))
	cashes := make([]int, len(values))
	copy(cashes, this.cashes)
	var x = amount
	for i := len(values) - 1; i >= 0; i-- {
		if cashes[i] == 0 {
			continue
		}
		if x >= values[i] {
			num := min(x/values[i], cashes[i])
			cashes[i] -= num
			ans[i] = num
			x -= num * values[i]
		}
	}
	if x != 0 {
		return []int{-1}
	}
	this.cashes = cashes
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	atm.Withdraw(600)
	atm.Deposit([]int{0, 1, 0, 1, 1})
	atm.Withdraw(600)
	atm.Withdraw(550)

}
