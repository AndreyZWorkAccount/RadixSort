package sortStructures

import "fmt"

//describes an element of a sorted array
type SortKey interface{
	fmt.Stringer

	IntKey() int
	Digit(exp int) int
}



