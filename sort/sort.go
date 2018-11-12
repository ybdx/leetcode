package sort

type Interface interface {
	Len() int
	Less() bool
}

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}
