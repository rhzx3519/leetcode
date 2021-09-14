package main

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	mapper := make(map[string]int)
	tables := make([]map[string]int, 501)
	for _, order := range orders {
		tableId, _ := strconv.Atoi(order[1])
		food := order[2]
		if tables[tableId] == nil {
			tables[tableId] = make(map[string]int)
		}
		tables[tableId][food]++
		mapper[food]++
	}

	foods := []string{}
	for food, _ := range mapper {
		foods = append(foods, food)
	}
	sort.Strings(foods)
	c := 1 + len(foods)
	ans := make([][]string, 1)
	ans[0] = []string{"Table"}
	ans[0] = append(ans[0], foods...)

	for i := 1; i < len(tables); i++ {
		if tables[i] == nil {
			continue
		}
		tmp := make([]string, c)
		tmp[0] = strconv.Itoa(i)
		for k := 1; k < c; k++ {
			food := ans[0][k]
			num := tables[i][food]
			tmp[k] = strconv.Itoa(num)
		}
		fmt.Println(tmp)
		ans = append(ans, tmp)
	}
	//fmt.Println(ans)

	return ans
}

func main() {
	displayTable([][]string{{"David","3","Ceviche"},{"Corina","10","Beef Burrito"},{"David","3","Fried Chicken"},{"Carla","5","Water"},{"Carla","5","Ceviche"},{"Rous","3","Ceviche"}})
}
