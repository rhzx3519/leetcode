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

	Comparator func(left, right T) int
)


var (
	// IntComparator is a Comparator for int
	IntComparator Comparator = func(left, right T) int {
		if left.(int) > right.(int) {
			return 1
		} else if left.(int) < right.(int) {
			return -1
		}
		return 0
	}

	// IntComparator is a Comparator for int64
	Int64Comparator Comparator = func(left, right T) int {
		if left.(int64) > right.(int64) {
			return 1
		} else if left.(int64) < right.(int64) {
			return -1
		}
		return 0
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
