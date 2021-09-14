package types

type (
	// T is a empty interface, that is `any` type.
	// since Go is not support generics now(but will coming soon),
	// so we use T to represent any type
	T interface {}

	K Comparable

	Comparable interface {
		CompareTo(o T) int
	}

	Comparator func(x, y T) int
)


var (
	// IntComparator is a Comparator for int
	IntComparator Comparator = func(x, y T) int {
		if x.(int) > y.(int) {
			return 1
		} else if x.(int) < y.(int) {
			return -1
		}
		return 0
	}

	// IntComparator is a Comparator for int64
	Int64Comparator Comparator = func(x, y T) int {
		if x.(int64) > y.(int64) {
			return 1
		} else if x.(int64) < y.(int64) {
			return -1
		}
		return 0
	}

	// return a reversed version of current Comparator function
	Reversed = func(cmp Comparator) Comparator {
		return func(x, y T) int {
			return -cmp(x, y)
		}
	}
)

type (
	ListNode struct {
		Value T
		Next  *ListNode
	}

	TreeNode struct {
		Value  T
		Parent *TreeNode
		Left   *TreeNode
		Right  *TreeNode
	}
)

// -------------------------------------------------------------------------------
// integer

type integer int

func (i integer) CompareTo(o T) int {
	j := o.(integer)
	if i > j {
		return 1
	} else if i < j {
		return -1
	}
	return 0
}
