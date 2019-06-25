package main

import (
	. "RadixSort/radixSort"
	. "RadixSort/sortStructures"
	"fmt"
)

func main() {

	arr := []SortKey{Int{55}, Int{17}, Int{33}, Int{2}, Int{1119}}

	var eng SortEngine = RadixSorter{}

	res := eng.DoSort(arr, Asc)

	for _, e := range res {
		fmt.Println(e.IntKey())
	}
}
