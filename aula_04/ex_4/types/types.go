package types

type Values []int

func (n *Values) Len() int {
	return len(*n)
}

func (n *Values) Less(i, j int) bool {
	return (*n)[i] < (*n)[j]
}

func (n *Values) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

type ValuesI interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

func NewValues(values Values) ValuesI {
	return &values
}
