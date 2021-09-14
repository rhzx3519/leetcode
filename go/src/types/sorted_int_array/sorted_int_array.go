/**
有序的int数组
 */
package sorted_int_array

/**
返回第一个arr中第一个大于等于k的元素下标，
如果k大于arr中的所有元素，则返回arr的长度
 */
func LowerBound(arr []int, k int) int {
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

/**
返回第一个arr中第一个大于的元素下标，
如果k大于arr中的所有元素，则返回arr的长度
 */
func UpperBound(arr []int, k int) int {
	l, r := 0, len(arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if arr[mid] > k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

/**
在i位置插入元素k,
i的范围是0~n
 */
func Insert(arr *[]int, i, k int) {
	n := len(*arr)
	if i < 0 || i > n {
		panic("Invalid insert position which should be in range [0, len(arr)].")
	}
	*arr = append(*arr, k)
	if i == len(*arr) {
		return
	}

	copy((*arr)[i+1:], (*arr)[i:len(*arr)-1])
	(*arr)[i] = k
}

// 返回含有重复元素的有序数组arr中，第一个等于k的元素下标，如果不存在则返回-1
func Find(arr []int, k int) int {
	i := LowerBound(arr, k)
	if i == len(arr) || arr[i] != k {
		return -1
	}
	return i
}

