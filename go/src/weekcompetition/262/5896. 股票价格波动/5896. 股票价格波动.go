/**
@author ZhengHao Lou
Date    2021/10/10
*/
package main

import "fmt"

type Stock struct {
	timestamp int
	price int
}

type StockPrice struct {
	stocks []Stock	// order by timestamp
	prices []int	// order by price
}

func Constructor() StockPrice {
	return StockPrice{
		stocks: []Stock{},
		prices: []int{},
	}
}


func (this *StockPrice) Update(timestamp int, price int)  {
	i := lowerBound(this.stocks, timestamp)
	if i != len(this.stocks) && this.stocks[i].timestamp == timestamp {

		j := lowerBoundInt(this.prices, this.stocks[i].price)	// delete
		copy(this.prices[j:], this.prices[j+1:])
		this.prices = this.prices[:len(this.prices) - 1]

		j = lowerBoundInt(this.prices, price)		// insert
		this.prices = append(this.prices, 0)
		copy(this.prices[j+1:], this.prices[j:])
		this.prices[j] = price

		this.stocks[i].price = price
		return
	}

	this.stocks = append(this.stocks, Stock{})
	copy(this.stocks[i+1:], this.stocks[i:])
	this.stocks[i] = Stock{
		timestamp: timestamp,
		price: price,
	}

	j := lowerBoundInt(this.prices, price)		// insert
	this.prices = append(this.prices, 0)
	copy(this.prices[j+1:], this.prices[j:])
	this.prices[j] = price
}


func (this *StockPrice) Current() int {
	fmt.Println(this.stocks[len(this.stocks) - 1].price)
	return this.stocks[len(this.stocks) - 1].price
}


func (this *StockPrice) Maximum() int {
	fmt.Println(this.prices)
	return this.prices[len(this.prices) - 1]
}


func (this *StockPrice) Minimum() int {
	fmt.Println(this.prices)
	return this.prices[0]
}

func lowerBound(stocks []Stock, timestamp int) int {
	l, r := 0, len(stocks)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if stocks[m].timestamp >= timestamp {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func lowerBoundInt(arr []int, k int) int {
	l, r := 0, len(arr)
	var mid int
	for l < r {
		mid = l + (r - l)>>1
		if arr[mid] >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	var stockPrice = Constructor();
	stockPrice.Update(1, 10); // 时间戳为 [1] ，对应的股票价格为 [10] 。
	stockPrice.Update(2, 5);  // 时间戳为 [1,2] ，对应的股票价格为 [10,5] 。
	stockPrice.Current();     // 返回 5 ，最新时间戳为 2 ，对应价格为 5 。
	stockPrice.Maximum();     // 返回 10 ，最高价格的时间戳为 1 ，价格为 10 。
	stockPrice.Update(1, 3);  // 之前时间戳为 1 的价格错误，价格更新为 3 。
	// 时间戳为 [1,2] ，对应股票价格为 [3,5] 。
	stockPrice.Maximum();     // 返回 5 ，更正后最高价格为 5 。
	stockPrice.Update(4, 2);  // 时间戳为 [1,2,4] ，对应价格为 [3,5,2] 。
	stockPrice.Minimum();     // 返回 2 ，最低价格时间戳为 4 ，价格为 2 。
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */