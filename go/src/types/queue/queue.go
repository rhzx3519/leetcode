package queue

type Queue struct {
	a []int
	cap int
	front, rear int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		a: make([]int, cap),
		cap: cap,
	}
}

func (q *Queue) Offer(x int) {
	if (q.rear + 1) % q.cap == q.front {
		panic("Queue is full.")
	}
	q.a[q.rear] = x
	q.rear = (q.rear + 1) % q.cap
}

func (q *Queue) Poll() (x int) {
	if q.isEmpty() {
		panic("Queue is empty.")
	}
	x = q.a[q.front]
	q.front = (q.front + 1) % q.cap
	return
}

func (q *Queue) isEmpty() bool {
	return q.front == q.rear
}

func (q *Queue) Size() int {
	if q.isEmpty() {
		return 0
	}
	if q.front < q.rear {
		return q.rear - q.front
	}
	return q.rear + q.cap - q.front
}


