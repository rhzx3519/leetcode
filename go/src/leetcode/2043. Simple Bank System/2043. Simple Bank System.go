/**
@author ZhengHao Lou
Date    2021/10/18
*/
package main

type Bank struct {
	balance []int64
	n int
}


func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
		n: len(balance),
	}
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.checkAccountId(account1) || !this.checkAccountId(account2) {
		return false
	}
	a, b := account1 - 1, account2 - 1
	if this.balance[a] < money {
		return false
	}
	this.balance[a] -= money
	this.balance[b] += money
	return true
}


func (this *Bank) Deposit(account int, money int64) bool {
	if !this.checkAccountId(account) {
		return false
	}
	this.balance[account - 1] += money
	return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.checkAccountId(account) {
		return false
	}
	if this.balance[account - 1] < money {
		return false
	}
	this.balance[account - 1] -= money
	return true
}

func (this *Bank) checkAccountId(account int) bool {
	return account >= 1 && account <= this.n
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
