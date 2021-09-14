package sorted_int_array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowerBound(t *testing.T) {
	arr := []int{1,1,4,7,10,10,50}
	n := len(arr)
	assert.Equal(t, n, LowerBound(arr, 51))
	assert.Equal(t, 0, LowerBound(arr, -1))
	assert.Equal(t, 0, LowerBound(arr, 1))
	assert.Equal(t, 0, LowerBound(arr, 0))
	assert.Equal(t, 2, LowerBound(arr, 2))
	assert.Equal(t, 2, LowerBound(arr, 3))
	assert.Equal(t, 2, LowerBound(arr, 4))
	assert.Equal(t, 3, LowerBound(arr, 5))
	assert.Equal(t, 4, LowerBound(arr, 8))
	assert.Equal(t, 4, LowerBound(arr, 10))
	assert.Equal(t, 6, LowerBound(arr, 11))
	assert.Equal(t, 6, LowerBound(arr, 50))
}

func TestUpperBound(t *testing.T) {
	arr := []int{1,1,4,7,7,7,7,10,10,50}
	//n := len(arr)
	assert.Equal(t, 0, UpperBound(arr, 0))
	assert.Equal(t, 2, UpperBound(arr, 1))
	assert.Equal(t, 2, UpperBound(arr, 2))
	assert.Equal(t, 3, UpperBound(arr, 4))
	assert.Equal(t, 3, UpperBound(arr, 6))
	assert.Equal(t, 7, UpperBound(arr, 7))
	assert.Equal(t, 9, UpperBound(arr, 10))
	assert.Equal(t, 9, UpperBound(arr, 49))
	assert.Equal(t, 10, UpperBound(arr, 50))
	assert.Equal(t, 10, UpperBound(arr, 51))
}

func TestInsert(t *testing.T) {
	arr := make([]int, 0)

	Insert(&arr, LowerBound(arr, 0), 0)
	assert.Equal(t, []int{0}, arr)

	Insert(&arr, LowerBound(arr, 1), 1)
	assert.Equal(t, []int{0,1}, arr)

	Insert(&arr, LowerBound(arr, 1), 1)
	assert.Equal(t, []int{0,1,1}, arr)

	Insert(&arr, LowerBound(arr, -5), -5)
	assert.Equal(t, []int{-5,0,1,1}, arr)

	Insert(&arr, LowerBound(arr, 5), 5)
	assert.Equal(t, []int{-5,0,1,1,5}, arr)

	Insert(&arr, LowerBound(arr, 100), 100)
	assert.Equal(t, []int{-5,0,1,1,5,100}, arr)

	Insert(&arr, LowerBound(arr, 5), 5)
	assert.Equal(t, []int{-5,0,1,1,5,5,100}, arr)

	Insert(&arr, LowerBound(arr, 3), 3)
	assert.Equal(t, []int{-5,0,1,1,3,5,5,100}, arr)

	Insert(&arr, LowerBound(arr, -10), -10)
	assert.Equal(t, []int{-10,-5,0,1,1,3,5,5,100}, arr)

	Insert(&arr, LowerBound(arr, -5), -5)
	assert.Equal(t, []int{-10,-5,-5,0,1,1,3,5,5,100}, arr)
}

func TestFind(t *testing.T) {
	arr := []int{1,1,4,7,7,7,7,10,10,50}

	assert.Equal(t, 0, Find(arr, 1))
	assert.Equal(t, 2, Find(arr, 4))
	assert.Equal(t, -1, Find(arr, 5))
	assert.Equal(t, 3, Find(arr, 7))
	assert.Equal(t, 7, Find(arr, 10))
	assert.Equal(t, -1, Find(arr, 11))
	assert.Equal(t, -1, Find(arr, 49))
	assert.Equal(t, 9, Find(arr, 50))
	assert.Equal(t, -1, Find(arr, 51))
}

// 维护一个有序的数组
func ExampleInsert() {
	arr := make([]int, 0)
	fmt.Println(arr)

	Insert(&arr, 0, 1)
	fmt.Println(arr)

	Insert(&arr, 0,1)
	fmt.Println(arr)

	Insert(&arr, 2,3)
	fmt.Println(arr)

	Insert(&arr, 2,2)
	fmt.Println(arr)

	// Output:
	// []
	// [1]
	// [1 1]
	// [1 1 3]
	// [1 1 2 3]

}


// 校验是否是非降序排列
func check_sorted(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			panic("check sorted failed.")
		}
	}
}
