package main

import (
	"sort"
	"strconv"
	"strings"
)

func countOfAtoms(formula string) string {
	n := len(formula)
	st := make([]int, 0)
	st2 := []map[string]int{}
	cur := make(map[string]int)

	var i, j, k, _ int
	for i < n {
		c := formula[i]
		if c == '(' {
			st = append(st, i)
			st2 = append(st2, cur)
			cur = make(map[string]int)
			i++
		} else if c == ')' {
			//last = st[len(st)-1]
			//st = st[:len(st)-1]
			j = i + 1
			for j < n && number(formula[j]) {
				j++
			}
			t := 1
			if j > i+1 {
				t, _ = strconv.Atoi(formula[i+1:j])
			}
			i = j

			for k, _ := range cur {
				cur[k] *= t
			}
			tmp := cur
			cur = st2[len(st2)-1]
			st2 = st2[:len(st2)-1]
			for k, v := range tmp {
				cur[k] += v
			}
		} else if upper(c) {
			j = i + 1
			for j < n && lower(formula[j]) {
				j++
			}
			name := formula[i:j]
			k = j
			for k < n && number(formula[k]) {
				k++
			}
			t := 1
			if k > j {
				t, _ = strconv.Atoi(formula[j:k])
			}
			j = k
			i = j

			cur[name] += t
		}

	}

	//fmt.Println(st2, cur)

	strs := []string{}
	keys := []string{}
	for k, _ := range cur {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		strs = append(strs, k)
		if cur[k] > 1 {
			strs = append(strs, strconv.Itoa(cur[k]))
		}
	}

	return strings.Join(strs, "")
}

func upper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func lower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func number(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	//countOfAtoms("H2O")
	countOfAtoms("Mg(OH)2")
	//countOfAtoms("K4(ON(SO3)2)2")
}