package main


type FindSumPairs struct {
	nums1, nums2 []int
	mapper map[int]int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	mapper := make(map[int]int)
	for i := range nums2 {
		mapper[nums2[i]]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, mapper: mapper}
}


func (this *FindSumPairs) Add(index int, val int)  {
	this.mapper[this.nums2[index]]--
	this.nums2[index] += val
	this.mapper[this.nums2[index]]++
}


func (this *FindSumPairs) Count(tot int) int {
	var count int
	for i := range this.nums1 {
		if n, ok := this.mapper[tot - this.nums1[i]]; ok {
			count += n
		}
	}
	return count
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */