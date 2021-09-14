package main

func subsetXORSum(nums []int) int {
	subsets := []int{0}
	for i := range nums {
		a := []int{}
		for _, sub := range subsets {
			a = append(a, sub^nums[i])
		}
		subsets = append(subsets, a...)
	}
	//fmt.Println(subsets)

	s := 0
	for _, num := range subsets {
		s += num
	}
	//fmt.Println(s)
	return s
}

func main() {
	subsetXORSum([]int{1,3})
	subsetXORSum([]int{5,1,6})
}
