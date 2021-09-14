package sorted_array

/**
返回第一个arr中第一个大于等于k的元素下标，
如果k大于arr中的所有元素，则返回arr的长度
*/
func LowerBound(arr []Comparable, k Comparable) int {
	l, r := 0, len(arr)
	var mid int
	for l < r {
		mid = l + (r - l)>>1
		if arr[mid].CompareTo(k) >= 0 {
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
func UpperBound(arr []Comparable, k Comparable) int {
	l, r := 0, len(arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if arr[mid].CompareTo(k) <= 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

/**
在i位置插入元素k,
i的范围是0~n
*/
func Insert(arr *[]Comparable, i int, k Comparable) {
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

// 删除下标i上的元素
func Delete(arr *[]Comparable, i int) {
	n := len(*arr)
	if i < 0 || i > n {
		panic("Invalid insert position which should be in range [0, len(arr)].")
	}
	copy((*arr)[i:n-1], (*arr)[i+1:])
	(*arr)[n-1] = nil
	*arr = (*arr)[:n-1]
}

// 返回含有重复元素的有序数组arr中，第一个等于k的元素下标，如果不存在则返回-1
func Find(arr []Comparable, k Comparable) int {
	i := LowerBound(arr, k)
	if i == len(arr) || arr[i] != k {
		return -1
	}
	return i
}


type (
	// T indicates any type
	T interface {}

	Comparable interface {
		CompareTo(o T) int
	}
)


type Integer int

func (i Integer) CompareTo(o T) int {
	var j = o.(Integer)
	if i > j {
		return 1
	} else if i < j {
		return -1
	} else {
		return 0
	}
}