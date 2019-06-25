package tests

import . "RadixSort/sortStructures"
import (
	"testing"
	"RadixSort/radixSort"
)


var testData = [][]SortKey{
	convert([]int{1,2,3,4,5,6,7,8,9}),
	convert([]int{11,2,3,4,54,6,7,8,93}),
	convert([]int{1,22,3,4,15,6,7,85,9}),
	convert([]int{999, 34, 5, 1011, 22, 16, 1, 77}),
	convert([]int{10,9,8,7,6,5,4,3,2,1}),
	convert([]int{1,2,3,4,5,3,7,8,9,11,2,3,4,54,6,7,8,93}),
	convert([]int{11,2,3,4,54,1,22,3,4,15,6,7,85,96,7,8,93}),
	convert([]int{1,22,3,4,15,6,7,85,9,7,6,5,4,3,2}),
	convert([]int{999, 34, 5, 1011, 22, 16, 1, 77}),
	convert([]int{10,9,8,7,6,5,4,3,2,1}),
}

func convert(arr []int) []SortKey{
	res := make([]SortKey,0)
	for _,e := range arr {
		res = append(res, Int{e})
	}
	return res
}

func TestSortingAsc(t *testing.T) {
	predicate := func(key1 SortKey, key2 SortKey) bool {
		return key1.IntKey() >= key2.IntKey()
	}
	testSorting(t, radixSort.Asc, predicate)
}

func TestSortingDesc(t *testing.T) {
	predicate := func(key1 SortKey, key2 SortKey) bool {
		return key1.IntKey() <= key2.IntKey()
	}
	testSorting(t, radixSort.Desc, predicate)
}

func testSorting(t *testing.T,direction radixSort.SortDirection, comparer func(key1 SortKey, key2 SortKey) bool ){
	var sorter radixSort.SortEngine = radixSort.RadixSorter{}
	for _,data := range testData{
		sortedArr := sorter.DoSort(data, direction)
		if !isSorted(sortedArr, comparer){
			t.Error("Should be sorted by ASC")
		}
	}
}

func isSorted(arr []SortKey, predicate func(key1 SortKey, key2 SortKey) bool) bool{
	for i := 1;i< len(arr);i++{
		if !predicate(arr[i],arr[i-1]){
			return false
		}
	}
	return true
}



