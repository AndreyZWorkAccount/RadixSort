package sortStructures

type Int struct{
	Val int
}

//SortKey<Int>
func (x Int) Digit(exp int) int{
	return (x.Val/exp)%10
}

func (x Int) IntKey() int{
	return x.Val
}

func (x Int) String() string{
	return string(x.Val)
}