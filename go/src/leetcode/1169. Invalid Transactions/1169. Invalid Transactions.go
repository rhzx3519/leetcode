package main

import (
	"fmt"
	"strconv"
	"strings"
)

type transaction struct {
	name string
	time int
	money int
	city string
	valid bool
}

func invalidTransactions(transactions []string) []string {
	var ans = []string{}

	mapper := make(map[string][]transaction)

	for _, t := range transactions {
		ts := strings.Split(t, ",")
		name := ts[0]
		time, _ := strconv.Atoi(ts[1])
		money, _ := strconv.Atoi(ts[2])
		city := ts[3]
		cur := transaction{name: name,
			time: time,
			money: money,
			city: city,
		valid: true}

		if money > 1000 {
			cur.valid = false
		}
		if _, ok := mapper[name]; !ok {
			mapper[name] = make([]transaction, 0)
		}

		for i, trans := range mapper[name] {
			if trans.city != city && abs(trans.time -  time) <= 60 {
				mapper[name][i].valid = false
				cur.valid = false
			}
		}

		mapper[name] = append(mapper[name], cur)
	}

	for _, ts := range mapper {
		for _, t := range ts {
			if !t.valid {
				s := strings.Join([]string{t.name, strconv.Itoa(t.time), strconv.Itoa(t.money), t.city}, ",")
				ans = append(ans, s)
			}
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return  x
}

func main() {
	fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","alice,50,100,beijing"}))
	fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","alice,50,1200,mtv"}))
	fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","bob,50,1200,mtv"}))
	fmt.Println(invalidTransactions([]string{"bob,689,1910,barcelona",
		"bob,832,1726,barcelona","bob,820,596,bangkok","bob,175,221,amsterdam"}))
}


