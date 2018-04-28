package algorithm

type queue struct {
	id []uint64
}

func (q queue) empty() bool {
	return len(q.id) == 0
}

func (q *queue) push(id uint64) {
	q.id = append(q.id, id)
}

func (q *queue) pop() uint64 {
	id := q.id[0]
	q.id = q.id[1:]
	return id
}
