package MultiSet

type MultiSet struct {
	arr []int
}

func NewMultiSet() *MultiSet {
	return &MultiSet{
		arr: make([]int, 0),
	}
}

func (m *MultiSet) Insert(x int) {
	i := m.UpperBound(x)
	m.arr = append(m.arr,  0)
	copy(m.arr[i+1:], m.arr[i:])
	m.arr[i] = x
}

func (m *MultiSet) LowBound(x int) int {
	l, r := 0, len(m.arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if m.arr[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func (m *MultiSet) UpperBound(x int) int {
	l, r := 0, len(m.arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if m.arr[mid] > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func (m *MultiSet) Erase(x int) {
	i := m.LowBound(x)
	if i == m.Size() || m.arr[i] != x {
		return
	}
	copy(m.arr[i:], m.arr[i+1:])
	m.arr = m.arr[:len(m.arr)-1]
}

func (m *MultiSet) RBegin() int {
	if m.IsEmpty() {
		panic("MultiSet is empty")
	}
	return m.arr[m.Size()-1]
}

func (m *MultiSet) Size() int {
	return len(m.arr)
}

func (m *MultiSet) IsEmpty() bool {
	return m.Size() == 0
}

