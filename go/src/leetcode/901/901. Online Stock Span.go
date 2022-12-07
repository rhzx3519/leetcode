/**
@author ZhengHao Lou
Date    2022/10/21
*/
package main

/**
https://leetcode.cn/problems/online-stock-span/
*/
type StockSpanner struct {
	spans  []int
	stocks []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	var w = 1
	for len(this.stocks) != 0 && this.stocks[len(this.stocks)-1] <= price {
		this.stocks = this.stocks[:len(this.stocks)-1]
		w += this.spans[len(this.spans)-1]
		this.spans = this.spans[:len(this.spans)-1]
	}
	this.stocks = append(this.stocks, price)
	this.spans = append(this.spans, w)
	return this.spans[len(this.spans)-1]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
