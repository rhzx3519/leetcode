package stack

type Stack []int

func NewStack() Stack {
	return []int{}
}

func (s Stack) Push(x int) {
	s = append(s, x)
}

func (s Stack) Pop() int {
	val := s[len(s)-1]
	s = s[:len(s)-1]
	return val
}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}
