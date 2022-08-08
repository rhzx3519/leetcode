/**
@author ZhengHao Lou
Date    2022/08/07
*/
package main

import "fmt"

const (
	equal2 = 1
	equal3 = 2
	inc3   = 3
)

func validPartition(nums []int) bool {
	n := len(nums)

	var f [][]int
	for i := 0; i < n; {
		var j = i + 1
		if j == n {
			return false
		}
		if j < n && nums[j] != nums[i] && nums[j] != nums[i]+1 {
			return false
		}
		if nums[j] == nums[i] {
			for j < n && nums[j] == nums[i] {
				j++
			}
			f = append(f, []int{j - i, 1, nums[i]})
		} else {
			for j < n && nums[j] != nums[j-1] {
				j++
			}
			f = append(f, []int{j - i, 2, nums[i]})
		}

		i = j
	}

	fmt.Println(f)
	for i, x := range f {
		if x[1] == 1 {
			if x[0]%2%3 != 0 && x[0]%3%2 != 0 {
				return false
			}
		} else {
			if x[0]%3 == 0 {
				continue
			}
			if (x[0]+1)%3 == 0 && i > 0 && f[i-1][1] == 1 && f[i-1][2] == (x[2]-1) {
				if (f[i-1][0]-1)%2%3 == 0 || (f[i-1][0]-1)%3%2 == 0 {
					continue
				}
			}
			return false
		}
	}

	return true
}

func main() {
	//fmt.Println(validPartition([]int{4, 4, 4, 5, 6}))
	//fmt.Println(validPartition([]int{4, 4, 4, 5, 6, 2}))
	//fmt.Println(validPartition([]int{4, 4, 1, 2, 3, 4}))
	//fmt.Println(validPartition([]int{803201, 803201, 803201, 803201, 803202, 803203}))

	fmt.Println(validPartition([]int{676575, 676575, 676575, 533985, 533985, 40495, 40495, 40495, 40495, 40495, 40495, 40495, 782020, 782021, 782022, 782023, 782024, 782025, 782026, 782027, 782028, 782029, 782030, 782031, 782032, 782033, 782034, 782035, 782036, 782037, 782038, 782039, 782040, 378070, 378070, 378070, 378071, 378072, 378073, 378074, 378075, 378076, 378077, 378078, 378079, 378080, 378081, 378082, 378083, 378084, 378085, 378086, 378087, 378088, 378089, 378090, 378091, 378092, 378093, 129959, 129959, 129959, 129959, 129959, 129959}))
}
