package main

import "fmt"

func getFolderNames(names []string) []string {
	var ans []string
	index := make(map[string]int) // 继续文件名为name的下一个最小整数序号
	for _, name := range names {
		if index[name] != 0 {
			k := index[name]
			for index[fmt.Sprintf("%v(%v)", name, k)] != 0 {
				k++
			}
			newName := fmt.Sprintf("%v(%v)", name, k)
			ans = append(ans, newName)
			index[newName]++
			index[name] = k + 1
		} else {
			ans = append(ans, name)
			index[name]++
		}
	}
	return ans
}

func main() {
	fmt.Println(getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)"}))
}
