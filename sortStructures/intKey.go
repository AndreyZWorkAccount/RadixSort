package sortStructures

type IntKeySource interface {
	intKey() int
}

type Int int

func (x *Int) intKey() int {
	return int(*x)
}
