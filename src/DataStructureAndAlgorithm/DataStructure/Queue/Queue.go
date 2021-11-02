package Queue

type Item interface {
}

type ItemQueue struct {
	items []Item
}

type ItemQueueer interface {
	New() *ItemQueue
	Enqueue(t Item)
	Dequeue() *Item
	IsEmpty() bool
	Size() int
}

func (s *ItemQueue) New() *ItemQueue {
	s.items = []Item{}
	return s
}

func (s *ItemQueue) Enqueue(t Item) {
	s.items = append(s.items, t)
}

func (s *ItemQueue) Dequeue() *Item {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return &item
}

func (s *ItemQueue) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *ItemQueue) Size() int {
	return len(s.items)
}
