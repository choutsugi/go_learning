package Stack

type Item struct {
}

type ItemStack struct {
	items []Item
}

type ItemStacker interface {
	New() *ItemStack
	Push(data Item) *ItemStack
	Pop() *Item
}

func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

func (s *ItemStack) Push(data Item) *ItemStack {
	s.items = append(s.items, data)
	return s
}

func (s *ItemStack) Pop() *Item {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return &item
}
