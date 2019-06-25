package radixSort

import . "RadixSort/sortStructures"

type SortDirection int

const(
	Asc SortDirection = iota
	Desc
)



type SortEngine interface {
	DoSort(arr []SortKey, direction SortDirection) []SortKey
}

//Implementation
type RadixSorter struct {
}

func (t RadixSorter) DoSort(arr []SortKey, direction SortDirection) []SortKey{
	resultArr := make([]SortKey, len(arr))
	copy(resultArr, arr)
	maxElem := findMaxElement(arr)
	for exp := 1; maxElem/exp > 0; exp *= 10{
		countSort(resultArr, exp, direction);
	}
	return resultArr
}

func countSort(arr []SortKey, exp int, direction SortDirection){
	arrayLen := len(arr)
	outArr := make([]SortKey, arrayLen)
	digitsCount := make([]int,10)

	for i := 0;i < arrayLen; i++{
		digitsCount[arr[i].Digit(exp)]++
	}

	updateByDirection(digitsCount, direction)

	for i := arrayLen-1;i>=0;i--{
		dgt := arr[i].Digit(exp)
		pos := digitsCount[dgt]-1
		outArr[pos] = arr[i]
		digitsCount[dgt]--
	}

	copy(arr,outArr)
}

func updateByDirection(digitsCount []int, direction SortDirection){
	switch direction {
	case Asc:
		for i := 1;i< 10;i++{
			digitsCount[i] += digitsCount[i-1]
		}
		break
	case Desc:
		for i := 8;i>=0;i--{
			digitsCount[i] += digitsCount[i+1]
		}
		break
	}
}

func findMaxElement(arr []SortKey) int{
	var ans int
	for i, key := range arr {
		e := key.IntKey()
		if i==0 || e > ans {
			ans = e
		}
	}
	return ans
}
