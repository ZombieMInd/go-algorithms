package queue

type Queue []interface{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}

	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue) Peek() interface{} {
	if len(*q) == 0 {
		return nil
	}

	return (*q)[0]
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Clear() {
	*q = (*q)[:0]
}
