package sorted_array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowerBound(t *testing.T) {
	arr := []Comparable{Integer(1),Integer(1),Integer(4),Integer(7),
		Integer(10),Integer(10),Integer(50)}
	n := len(arr)
	assert.Equal(t, n, LowerBound(arr, Integer(51)))
	assert.Equal(t, 0, LowerBound(arr, Integer(-1)))
	assert.Equal(t, 0, LowerBound(arr, Integer(1)))
	assert.Equal(t, 0, LowerBound(arr, Integer(0)))
	assert.Equal(t, 2, LowerBound(arr, Integer(2)))
	assert.Equal(t, 2, LowerBound(arr, Integer(3)))
	assert.Equal(t, 2, LowerBound(arr, Integer(4)))
	assert.Equal(t, 3, LowerBound(arr, Integer(5)))
	assert.Equal(t, 4, LowerBound(arr, Integer(8)))
	assert.Equal(t, 4, LowerBound(arr, Integer(10)))
	assert.Equal(t, 6, LowerBound(arr, Integer(11)))
	assert.Equal(t, 6, LowerBound(arr, Integer(50)))
}

func TestUpperBound(t *testing.T) {
	arr := []Comparable{Integer(1),Integer(1),Integer(4),Integer(7),Integer(7),Integer(7),Integer(7),Integer(10),Integer(10),Integer(50)}
	//n := len(arr)
	assert.Equal(t, 0, UpperBound(arr, Integer(0)))
	assert.Equal(t, 2, UpperBound(arr, Integer(1)))
	assert.Equal(t, 2, UpperBound(arr, Integer(2)))
	assert.Equal(t, 3, UpperBound(arr, Integer(4)))
	assert.Equal(t, 3, UpperBound(arr, Integer(6)))
	assert.Equal(t, 7, UpperBound(arr, Integer(7)))
	assert.Equal(t, 9, UpperBound(arr, Integer(10)))
	assert.Equal(t, 9, UpperBound(arr, Integer(49)))
	assert.Equal(t, 10, UpperBound(arr, Integer(50)))
	assert.Equal(t, 10, UpperBound(arr, Integer(51)))
}

func TestInsert(t *testing.T) {
	arr := make([]Comparable, 0)

	Insert(&arr, LowerBound(arr, Integer(0)), Integer(0))
	assert.Equal(t, []Comparable{Integer(0)}, arr)

	Insert(&arr, LowerBound(arr, Integer(1)), Integer(1))
	assert.Equal(t, []Comparable{Integer(0),Integer(1)}, arr)

	Insert(&arr, LowerBound(arr, Integer(1)), Integer(1))
	assert.Equal(t, []Comparable{Integer(0),Integer(1),Integer(1)}, arr)

	Insert(&arr, LowerBound(arr, Integer(-5)), Integer(-5))
	assert.Equal(t, []Comparable{Integer(-5),Integer(0),Integer(1),Integer(1)}, arr)

	Insert(&arr, LowerBound(arr, Integer(5)), Integer(5))
	assert.Equal(t, []Comparable{Integer(-5),Integer(0),Integer(1),Integer(1),Integer(5)}, arr)

	Insert(&arr, LowerBound(arr, Integer(100)), Integer(100))
	assert.Equal(t, []Comparable{Integer(-5),Integer(0),Integer(1),Integer(1),Integer(5),Integer(100)}, arr)

	Insert(&arr, LowerBound(arr, Integer(5)), Integer(5))
	assert.Equal(t, []Comparable{Integer(-5),Integer(0),Integer(1),Integer(1),Integer(5),Integer(5),Integer(100)}, arr)

	Insert(&arr, LowerBound(arr, Integer(3)), Integer(3))
	assert.Equal(t, []Comparable{Integer(-5),Integer(0),Integer(1),Integer(1),Integer(3),Integer(5),Integer(5),Integer(100)}, arr)

	Insert(&arr, LowerBound(arr, Integer(-10)), Integer(-10))
	assert.Equal(t, []Comparable{Integer(-10),Integer(-5),Integer(0),Integer(1),Integer(1),Integer(3),Integer(5),Integer(5),Integer(100)}, arr)

	Insert(&arr, LowerBound(arr, Integer(-5)), Integer(-5))
	assert.Equal(t, []Comparable{Integer(-10),Integer(-5),Integer(-5),Integer(0),Integer(1),Integer(1),Integer(3),Integer(5),Integer(5),Integer(100)}, arr)
}

func TestDelete(t *testing.T) {
	arr := []Comparable{Integer(1),Integer(1),Integer(4),Integer(7),Integer(7),Integer(7),Integer(7),Integer(10),Integer(10),Integer(50)}

	Delete(&arr, 0)
	assert.Equal(t, []Comparable{Integer(1),Integer(4),Integer(7),Integer(7),Integer(7),Integer(7),Integer(10),Integer(10),Integer(50)}, arr)

	Delete(&arr, 3)
	assert.Equal(t, []Comparable{Integer(1),Integer(4),Integer(7),Integer(7),Integer(7),Integer(10),Integer(10),Integer(50)}, arr)


}

func TestFind(t *testing.T) {
	arr := []Comparable{Integer(1),Integer(1),Integer(4),Integer(7),Integer(7),Integer(7),Integer(7),Integer(10),Integer(10),Integer(50)}

	assert.Equal(t, 0, Find(arr, Integer(1)))
	assert.Equal(t, 2, Find(arr, Integer(4)))
	assert.Equal(t, -1, Find(arr, Integer(5)))
	assert.Equal(t, 3, Find(arr, Integer(7)))
	assert.Equal(t, 7, Find(arr, Integer(10)))
	assert.Equal(t, -1, Find(arr, Integer(11)))
	assert.Equal(t, -1, Find(arr, Integer(49)))
	assert.Equal(t, 9, Find(arr, Integer(50)))
	assert.Equal(t, -1, Find(arr, Integer(51)))
}

// 维护一个有序的数组
func ExampleInsert() {
	arr := make([]Comparable, 0)
	fmt.Println(arr)

	Insert(&arr, 0, Integer(1))
	fmt.Println(arr)

	Insert(&arr, 0,Integer(1))
	fmt.Println(arr)

	Insert(&arr, 2,Integer(3))
	fmt.Println(arr)

	Insert(&arr, 2,Integer(2))
	fmt.Println(arr)

	// Output:
	// []
	// [1]
	// [1 1]
	// [1 1 3]
	// [1 1 2 3]

}

