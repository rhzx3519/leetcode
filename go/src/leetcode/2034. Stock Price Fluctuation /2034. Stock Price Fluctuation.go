/**
@author ZhengHao Lou
Date    2022/01/23
*/
package main

type quote struct {
	price, timestamp int
}

type StockPrice struct {
	quotes []quote
	prices []int
}

func Constructor() StockPrice {
	return StockPrice{
		quotes: []quote{},
		prices: []int{},
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	i := lowerBoundQuotes(this.quotes, timestamp)
	if i == len(this.quotes) || this.quotes[i].timestamp != timestamp { // new quote
		this.quotes = append(this.quotes, quote{})
		copy(this.quotes[i+1:], this.quotes[i:])
		this.quotes[i] = quote{
			timestamp: timestamp,
			price:     price,
		}
		j := lowerBound(this.prices, price)
		this.prices = append(this.prices, 0)
		copy(this.prices[j+1:], this.prices[j:])
		this.prices[j] = price
		return
	}
	// update quote
	j := lowerBound(this.prices, this.quotes[i].price)
	copy(this.prices[j:], this.prices[j+1:]) // delete wrong price
	this.prices = this.prices[:len(this.prices)-1]
	j = lowerBound(this.prices, price)
	this.prices = append(this.prices, 0)
	copy(this.prices[j+1:], this.prices[j:]) // insert right price
	this.prices[j] = price

	this.quotes[i].price = price
}

func (this *StockPrice) Current() int {
	return this.quotes[len(this.quotes)-1].price
}

func (this *StockPrice) Maximum() int {
	return this.prices[len(this.prices)-1]
}

func (this *StockPrice) Minimum() int {
	return this.prices[0]
}

func lowerBoundQuotes(quotes []quote, x int) int {
	l, r := 0, len(quotes)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if quotes[m].timestamp >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
